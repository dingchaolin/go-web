package main

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"

)

func main() {

	//beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{})
	//自动路由
//localhost:8080/main/get
//localhost:8080/main/add 这样访问即可
//使用自动路由必须路由结构的名称以controller结尾才行
	beego.AutoRouter(&controllers.MainController{})

	beego.Run()
}

