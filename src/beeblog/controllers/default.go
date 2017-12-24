package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"

	c.Data["TrueCond"] = true;
	c.Data["FalseCond"] = false;

	type u struct{
		Name string
		Age int
		Sex string
		School string
	}

	user := &u{
		Name: "dcl",
		Age : 20,
		Sex : "Male",
		School: "THU",
	}

	c.Data["User"] = user

	nums := []int{1,2,3,4}

	c.Data["Nums"] = nums


	//模板变量
	c.Data["TplVar"] = "hello guys"

	//将字符串转化为html
	c.Data["Html"] = "<div> Hello world</div>"

	c.Data["Pipe"] = "<div> Hello Pipe</div>"
}
//controller 接口返回多个值 数据 和 error
func (c *MainController) Post() {

	//cookie验证
	if !checkAccount( c.Ctx ){
		c.Redirect("/login", 302)
		return
	}
	name := c.Input().Get("name")


	fmt.Println( name )
	maxAge := 0
	if name != "" {
		maxAge = 60
	}
	c.Ctx.SetCookie("name", name, maxAge, "/")

	c.Ctx.WriteString("ok")
	return
}

func checkAccount( c *context.Context) bool{
	_, err := c.Request.Cookie("name")
	if err != nil{
		return  true
	}else{
		return false
	}


}

func (c *MainController) Add() {

	c.Ctx.WriteString("Add")
	return
}

func (c *MainController) Find() {
// restful路由参数的获取
    param0 := c.Ctx.Input.Param( "0")
	param1 := c.Ctx.Input.Param( "1")
	c.Ctx.WriteString(param0 + "-" + param1)
	return
}