package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"imooc-iris/common"
	"imooc-iris/datamodels"
	"imooc-iris/services"
)

type ProductController struct {
	Ctx iris.Context // 上下文
	ProductService services.IProductService
}

func (p *ProductController) GetAll() mvc.View{
	productArray,_ := p.ProductService.GetAllProduct()
	return mvc.View{
		Name:"product/view.html",
		Data:iris.Map{
			"productArray":productArray,
		},
	}
}

// 修改商品
func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName:"imooc"})
	if err := dec.Decode(p.Ctx.Request().Form,product);err !=nil{
		p.Ctx.Application().Logger().Debug()
	}
	err:= p.ProductService.UpdateProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug()
	}
	p.Ctx.Redirect("/product/all")
}