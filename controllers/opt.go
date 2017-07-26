package controllers

import (
	"dnsadmin/models"
	"fmt"

	"github.com/astaxie/beego"
)

type EnableController struct {
	beego.Controller
}

func (c *EnableController) Post() {
	arg := c.Ctx.Input.Param(":splat")
	var rec models.Records
	rec.Name = arg
	rec.Modifier_ip = c.Ctx.Request.RemoteAddr
	err := models.Enable(rec)
	var resp Dto
	if err == nil {
		resp.Code = "ok"
		resp.Msg = fmt.Sprintf("enable %v success", rec.Name)
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		resp.Code = "error"
		resp.Msg = fmt.Sprintf("enable %v fail, with error %v", rec.Name, err.Error())
		c.Data["json"] = resp
		c.ServeJSON()
	}
}

type DisableController struct {
	beego.Controller
}

func (c *DisableController) Post() {
	arg := c.Ctx.Input.Param(":splat")
	var rec models.Records
	rec.Name = arg
	rec.Modifier_ip = c.Ctx.Request.RemoteAddr
	err := models.Disable(rec)
	var resp Dto
	if err == nil {
		resp.Code = "ok"
		resp.Msg = fmt.Sprintf("disable %v success", rec.Name)
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		resp.Code = "error"
		resp.Msg = fmt.Sprintf("disable %v fail, with error %v", rec.Name, err.Error())
		c.Data["json"] = resp
		c.ServeJSON()
	}
}

type SyncAllController struct {
	beego.Controller
}

func (c *SyncAllController) Post() {
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
