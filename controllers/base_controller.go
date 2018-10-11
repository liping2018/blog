package controllers

import (
	"my_blog/utils"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin int //用户是否已登录
}

func init() {

}

func (c *BaseController) ok() {
	//models.AddNewUser("liping", "12345", "你好啊", "www.baidu.com", 1)
	utils.SetRedisValue("liping", "110", "50")
	r := &utils.R{0, "Success", make(map[string]interface{}, 0)}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

//平台统一自定义异常
func (c *BaseController) RRException(msg string) {
	r := &utils.R{500, msg, make(map[string]interface{})}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()

}
