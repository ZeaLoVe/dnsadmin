package controllers

import (
	"dnsadmin/models"
	"fmt"
	"github.com/astaxie/beego"
)

type AddController struct {
	beego.Controller
}

func (c *AddController) Get() {
	c.Data["Website"] = "DNSadmin"
	c.TplNames = "add.tpl"
}

type InsertController struct {
	beego.Controller
}

func (c *InsertController) Post() {
	domain := c.GetString("domain")
	content := c.GetString("content")
	ttl, err := c.GetInt("ttl")
	auth := c.GetString("auth")
	if err != nil {
		ttl = 0
	}
	rec := models.Records{}
	rec.Name = domain
	rec.Content = content
	rec.Ttl = ttl
	rec.Modifier_ip = c.Ctx.Request.RemoteAddr
	if auth == "" {
		rec.Auth = c.Ctx.Request.RemoteAddr
	} else {
		rec.Auth = auth
	}
	err = models.Save(rec)
	if err == nil {
		c.Ctx.Redirect(302, "/")
	} else {
		msg := fmt.Sprintf("Fail to save info in database with err:%v", err.Error())
		c.Ctx.WriteString(msg)
	}

}
