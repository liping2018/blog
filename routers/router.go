package routers

import (
	"my_blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{})
	beego.Router("/login", &controllers.UserController{}, "get,post:Login")
	//添加注解路由
	ns := beego.NewNamespace("v1",
		beego.NSNamespace("test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
