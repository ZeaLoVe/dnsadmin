package controllers

import (
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
