package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["my_blog/controllers:TestController"] = append(beego.GlobalControllerRouter["my_blog/controllers:TestController"],
		beego.ControllerComments{
			Method: "TestRequestMessage",
			Router: `test`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
