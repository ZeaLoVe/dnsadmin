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
	c.Data["Website"] = "DNSadmin"
	arg := c.Ctx.Input.Param(":splat")
	var rec models.Records
	rec.Name = arg
	rec.Modifier_ip = c.Ctx.Request.RemoteAddr
	err := models.Enable(rec)
	if err == nil {
		c.Ctx.Redirect(302, "/")
	} else {
		msg := fmt.Sprintf("Fail to enable domain %v with err:%v", arg, err.Error())
		c.Ctx.WriteString(msg)
	}
}

type DisableController struct {
	beego.Controller
}

func (c *DisableController) Get() {
	c.Data["Website"] = "DNSadmin"
	arg := c.Ctx.Input.Param(":splat")
	var rec models.Records
	rec.Name = arg
	rec.Modifier_ip = c.Ctx.Request.RemoteAddr
	err := models.Disable(rec)
	if err == nil {
		c.Ctx.Redirect(302, "/")
	} else {
		msg := fmt.Sprintf("Fail to disable domain %v with err:%v", arg, err.Error())
		c.Ctx.WriteString(msg)
	}
}

type SyncAllController struct {
	beego.Controller
}

func (c *SyncAllController) Get() {
	num, err := models.SyncAll()
	var msg string
	if err == nil {
		msg = fmt.Sprint("SyncAll records %v all success", num)
	} else {
		msg = fmt.Sprintf("SyncAll records with success number:%v.\nError List:%v", num, err.Error())
	}
	c.Ctx.WriteString(msg)
}
