package controllers

import (
	"dnsadmin/models"
	"fmt"
	"github.com/astaxie/beego"
)

type EnableController struct {
	beego.Controller
}

func (c *EnableController) Get() {
	sess := c.StartSession()
	username := sess.Get("user_name")
	if username == "" {
		c.Ctx.Redirect(302, "/")
		return
	}
	c.Data["UserName"] = username
	c.Data["Website"] = "DNSadmin"
	arg := c.Ctx.Input.Param(":splat")
	var rec models.Records
	rec.Name = arg
	rec.Modifier_ip = c.Ctx.Request.RemoteAddr
	err := models.Enable(rec)
	if err == nil {
		beego.BeeLogger.Info("user:%v enable a domain< %v > success", username, rec.Name)
		c.Ctx.Redirect(302, "/")
	} else {
		beego.BeeLogger.Info("user:%v add a domain< %v > fail with err:%v", username, rec.Name, err.Error())
		msg := fmt.Sprintf("Fail to enable domain %v with err:%v", arg, err.Error())
		c.Ctx.WriteString(msg)
	}
}

type DisableController struct {
	beego.Controller
}

func (c *DisableController) Get() {
	sess := c.StartSession()
	username := sess.Get("user_name")
	if username == "" {
		c.Ctx.Redirect(302, "/")
		return
	}
	c.Data["UserName"] = username
	c.Data["Website"] = "DNSadmin"
	arg := c.Ctx.Input.Param(":splat")
	var rec models.Records
	rec.Name = arg
	rec.Modifier_ip = c.Ctx.Request.RemoteAddr
	err := models.Disable(rec)
	if err == nil {
		beego.BeeLogger.Info("user:%v disable a domain< %v > success", username, rec.Name)
		c.Ctx.Redirect(302, "/")
	} else {
		beego.BeeLogger.Info("user:%v disable a domain< %v > fail with err:%v", username, rec.Name, err.Error())
		msg := fmt.Sprintf("Fail to disable domain %v with err:%v", arg, err.Error())
		c.Ctx.WriteString(msg)
	}
}

type SyncAllController struct {
	beego.Controller
}

func (c *SyncAllController) Get() {
	sess := c.StartSession()
	username := sess.Get("user_name")
	if username == "" {
		c.Ctx.Redirect(302, "/")
		return
	}
	num, err := models.SyncAll()
	var msg string
	if err == nil {
		beego.BeeLogger.Info("user:%v sync all domain success", username)
		msg = fmt.Sprintf("SyncAll records %v all success", num)
	} else {
		beego.BeeLogger.Info("user:%v sync all domain fail with err:%v", username, err.Error())
		msg = fmt.Sprintf("SyncAll records with success number:%v.\nError List:%v", num, err.Error())
	}
	c.Ctx.WriteString(msg)
}
