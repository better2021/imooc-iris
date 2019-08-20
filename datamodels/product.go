package datamodels

type Product struct {
	ID int64 `json:"id" sql:"ID" imooc:"id"`
	ProductName string `json:"ProductName" sql:"ProductName" imooc:"ProductName"`
	ProductNum int64 `json:"ProductNum" sql:"ProductNum" imooc:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"ProductImage" imooc:"ProductImage"`
	ProductUrl string `json:"productUrl" sql:"ProductUrl" imooc:"ProductUrl"`
}
