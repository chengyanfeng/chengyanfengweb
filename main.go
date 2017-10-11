package main

import (
	_ "chengyanfengweb/routers"
	"github.com/astaxie/beego"
	."chengyanfengweb/models"
	_"chengyanfengweb/controllers"
	//"chengyanfengweb/controllers"
	"chengyanfengweb/controllers"
)
//创建一个Gorm 的类型
var gorm Gorm
func main() {
	//数据库链接和数据表创建初始化
	gorm.Init()
	//端口
	beego.BConfig.Listen.HTTPPort = 8080                     //端口设置
	beego.BConfig.RecoverPanic = true                        //开启异常捕获
	beego.BConfig.EnableErrorsShow = true
	beego.BConfig.CopyRequestBody = true

	//自动注册路由
	beego.AutoRouter(&controllers.ChengController{})
	beego.Run()
}

