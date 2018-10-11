//file:sysinit.go
//auth:liping
//date:2018-09-25
//desc

package sysinit

//系统初始化
func init() {
	//初始化数据库
	InitDatabase()
	//初始化redis
	//utils.InitRedis()
}
