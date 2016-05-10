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
	sess := c.StartSession()
	username := sess.Get("user_name")
	if username == "" {
		c.Ctx.Redirect(302, "/")
		return
	}
	c.Data["UserName"] = username
	c.Data["Website"] = "DNSadmin"
	c.TplNames = "add.tpl"
}

type InsertController struct {
	beego.Controller
}

func (c *InsertController) Post() {
	sess := c.StartSession()
	username := sess.Get("user_name")
	if username == "" {
		c.Ctx.Redirect(302, "/")
		return
	}
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
		beego.BeeLogger.Info("user:%v add a domain< %v > success", username, domain)
		c.Ctx.Redirect(302, "/")
	} else {
		beego.BeeLogger.Info("user:%v add a domain< %v > fail with err:%v", username, domain, err.Error())
		msg := fmt.Sprintf("Fail to save info in database with err:%v", err.Error())
		c.Ctx.WriteString(msg)
	}

}
