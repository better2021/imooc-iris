package main

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"imooc-iris/backend/web/controllers"
	"imooc-iris/common"
	"imooc-iris/repositories"
	"imooc-iris/services"
)

func main(){
	// 1.创建iris实例
	app := iris.New()
	// 2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	// 3.注册模板
	template := iris.HTML("./backend/web/views",".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	// 4.设置模板目录
	app.StaticWeb("/assets","./backend/web/assets")
	// 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",ctx.Values().GetStringDefault("message","访问的页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	// 连接数据库
	db,err := common.NewMysqlConn()
	if err !=nil{
		fmt.Println(err)
	}
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	// 5.注册控制器
	productRepository := repositories.NewProductManager("product",db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx,productService)
	product.Handle(new(controllers.ProductController))

	// 6.启动服务
	app.Run(
		iris.Addr("location:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}
