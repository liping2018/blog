//file:filter.go
//auth:liping
//date:2018-09-25
//desc:
package filter

import (
	"my_blog/utils"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
)

var urls = []string{
	"/",
	"/login",
	"/v1/test/test",
	"/static/",
}

func Ignoretoken(contexUri string) bool {
	for _, uri := range urls {
		if strings.Contains(contexUri, uri) {
			return true
		}
	}
	return false
}

var userFilter = func(ctx *context.Context) {

	//判断是否忽略过滤
	if !Ignoretoken(ctx.Request.RequestURI) {
		//获取session
		user := ctx.Input.Session("user")
		if user == nil {
			//未登录
			r := &utils.R{utils.USER_UNLOGIN_CODE, utils.USER_UNLOGIN_MSG, make(map[string]string)}
			ctx.Output.JSON(r, true, true)
		}
	}
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, userFilter)
}
