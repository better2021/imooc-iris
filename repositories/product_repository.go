package repositories

import (
	"database/sql"
	"imooc-iris/common"
	"imooc-iris/datamodels"
	"strconv"
)

// 第一步，先开发对应的接口
type IProduct interface {
	// 连接数据库
	Conn()(error)
	Insert(*datamodels.Product)(int64,error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64)(*datamodels.Product,error)
	SelectAll()([]*datamodels.Product,error)
}

//第二步，实现定义的接口
type ProductManager struct {
	table string
	mysqlConn *sql.DB
}

func NewProductManager(table string,db *sql.DB) IProduct{
	return &ProductManager{table:table,mysqlConn:db}
}

// 数据库连接
func (p *ProductManager) Conn()(err error){
	if p.mysqlConn ==nil{
		mysql,err := common.NewMysqlConn()
		if err !=nil{
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == ""{
		p.table = "product"
	}
	return
}

// 数据库插入
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64,err error){
	// 1.判断连接是否存在
	if err = p.Conn();err != nil {
		return
	}
	// 2.准备sql
	sql := "INSERT product SET productName=?,productNum=?,productImage=?,productUrl=?"
	stmt,errSql := p.mysqlConn.Prepare(sql)
	if errSql !=nil{
		return 0,errSql
	}
	// 3.传入参数
	result,errStmt := stmt.Exec(product.ProductName,product.ProductNum,product.ProductUrl,product.ProductImage)
	if errStmt !=nil{
		return 0,errStmt
	}
	return result.LastInsertId()
}

// 商品删除
func (p *ProductManager) Delete(prodictID int64) bool{
	// 1.判断连接是否存在
	if err := p.Conn();err!=nil{
		return false
	}
	sql := "delete from product where ID=?"
	stmt,err := p.mysqlConn.Prepare(sql)
	if err ==nil{
		return false
	}
	_,err = stmt.Exec(prodictID)
	if err !=nil{
		return false
	}
	return true
}

// 商品更新
func (p *ProductManager) Update(product *datamodels.Product) error{
	// 1.判断连接是否存在
	if err := p.Conn();err !=nil{
		return err
	}
	sql := "Update product set productName=?,productNum=?,productImage=?,productUrl=? where ID=" + strconv.FormatInt(product.ID,10)

	stmt,err := p.mysqlConn.Prepare(sql)
	if err !=nil{
		return  err
	}

	_,err = stmt.Exec(product.ProductName,product.ProductNum,product.ProductUrl,product.ProductImage)
	if err !=nil{
		return err
	}
	return nil
}

// 根据商品ID查询商品
func (p *ProductManager) SelectByKey(productID int64)(product *datamodels.Product,err error) {
	// 1.判断连接是否存在
	if err = p.Conn();err !=nil{
		return &datamodels.Product{},err
	}
	sql := "Select * from " + p.table + "where ID=" + strconv.FormatInt(productID,10)
	row,errRow := p.mysqlConn.Query(sql)
	defer row.Close()
	if errRow !=nil{
		return &datamodels.Product{},errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{},nil
	}
	common.DataToStructByTagSql(result,product)
	return
}

// 获取所有商品
func (p *ProductManager) SelectAll()(productArray []*datamodels.Product,errProduct error){
	// 判断连接是否存在
	if err := p.Conn();err!=nil{
		return nil,err
	}
	sql := "Select * from " +p.table
	rows,err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err!=nil{
		return nil,err
	}
	result := common.GetResultRows()
	if len(result)==0{
		return nil,nil
	}
	for _,v:= range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v,product)
		productArray = append(productArray, product)
	}
	return
}