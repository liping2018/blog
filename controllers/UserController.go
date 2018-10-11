//file:UserController.go
//auth:lushifu
//date:2018-09-26
//desc:

package controllers

import (
	"fmt"
	"my_blog/service"
	"my_blog/utils"
	"strconv"

	"github.com/astaxie/beego/validation"
)

type UserController struct {
	BaseController
}

type UserLogin struct {
	Username string `valid:"Required"`
	Password string `valid:"Required"`
}

func (c *UserController) Get() {
	c.TplName = "login.html"
}

func (c *UserController) Login() {

	if c.Ctx.Request.Method == "GET" {
		c.TplName = "login.html"
		return
	}
	var username, password string
	var remenber int
	username = c.GetString("Username")
	password = c.GetString("Password")
	remenber, _ = c.GetInt("remenber")
	fmt.Println(username, password, remenber)
	validate := validation.Validation{}
	userlogin := UserLogin{Username: username, Password: password}
	b, err := validate.Valid(&userlogin)
	if err != nil {
		c.RRException("验证用户名密码错误")
	}
	if !b {
		c.RRException("用户名或密码格式错误")
	}
	user := service.QueryUserInfoByUername(username)
	depassword := utils.MD5Encode(password)
	if depassword == user.Password {
		c.SetSession("user", user)
		RememberUserNameAndPwd(c, remenber, username, password) //记住用户名密码
		c.Data["actionName"] = "home"
		c.TplName = "blog/home.html"
	} else {
		c.RRException("用户名或密码不匹配")
	}
}

//记住密码
func RememberUserNameAndPwd(c *UserController, remeber int, username, password string) {
	if remeber == 1 {
		lifetime := utils.GetString("cookie::lifetime")
		lifet, err := strconv.Atoi(lifetime)
		if err != nil {
			utils.Expel(utils.UNNKOW_ERROR_CODE, utils.UNNKOW_ERROR_MSG)
		}
		fmt.Println("lifetime", lifetime)
		c.Ctx.SetCookie("loginInfo", username+","+password, lifet, "/")
	}
}
