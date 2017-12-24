package main

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"

)

func main() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{})

	beego.Run()
}

