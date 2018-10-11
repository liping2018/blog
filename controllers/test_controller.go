package controllers

import "github.com/astaxie/beego"

type TestController struct {
	BaseController
}

func (c *TestController) URLMapping() {
	c.Mapping("test", c.TestRequestMessage)
}

//@router test [get]
func (c *TestController) TestRequestMessage() {
	flash := beego.NewFlash()
	flash.Success("访问成功了呢")
	flash.Store(&c.Controller)
	c.Data["liping"] = "李平"
	c.TplName = "index.tpl"

}
