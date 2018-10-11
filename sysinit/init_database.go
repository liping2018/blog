//file:initDatabase.go
//auth:liping
//date:2018-09-24
//desc:初始化数据库工具

package sysinit

import (
	"my_blog/models"
	"my_blog/utils"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() {
	database := utils.GetString("mysql::database")
	username := utils.GetString("mysql::username")
	password := utils.GetString("mysql::password")
	host := utils.GetString("mysql::host")
	port := utils.GetString("mysql::port")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", username+":"+password+"@("+host+":"+port+")/"+database+"?charset=utf8")
	orm.RegisterModel(
		new(models.Message),
		new(models.User),
		new(models.Articel),
		new(models.Images),
		new(models.Reply),
		new(models.Type))
	orm.RunSyncdb("default", false, true)
}
