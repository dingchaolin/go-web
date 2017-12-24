package controllers

import (
	"github.com/astaxie/beego"
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

func (c *MainController) Post() {
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

func checkAccount( c *beego.Controller) bool{
	_, err := c.Ctx.Request.Cookie("name")
	if err != nil{
		return  true
	}else{
		return false
	}


}