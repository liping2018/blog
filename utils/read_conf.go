//file:read_conf.go
//auth:liping
//date:2018-09-24
//desc:读取conf配置文件

package utils

import (
	"fmt"

	"github.com/astaxie/beego"
)

func GetString(key string) string {
	fmt.Println(beego.AppConfig.String(key))
	return beego.AppConfig.String(key)
}
