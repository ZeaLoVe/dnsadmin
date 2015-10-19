package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "DNSadmin"
	c.Layout = "index.tpl"
	c.TplNames = "data.tpl"
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Ctx.WriteString(fmt.Sprint(c.GetStrings("Counter")))
	c.Ctx.WriteString(fmt.Sprint(c.GetString("Endpoint")))
}
