package controllers

import (
	"dnsadmin/models"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

type Dto struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

type DomainsController struct {
	beego.Controller
}

//获取域名
func (c *DomainsController) Get() {
	var res []models.Records
	name := c.GetString("name")
	content := c.GetString("content")
	auth := c.GetString("auth")

	sql := "select * from records"
	seq := "%"
	var filters []string
	if name != "" {
		filter := fmt.Sprintf("name like '%v%v%v'", seq, name, seq)
		filters = append(filters, filter)
	}
	if content != "" {
		filter := fmt.Sprintf("content like '%v%v%v'", seq, content, seq)
		filters = append(filters, filter)
	}
	if auth != "" {
		filter := fmt.Sprintf("auth like '%v%v%v'", seq, auth, seq)
		filters = append(filters, filter)
	}
	for i, filter := range filters {
		if i == 0 {
			sql += " where " + filter
		} else {
			sql += " and " + filter
		}
	}
	count_sql := strings.Replace(sql, "*", "count(*)", 1)
	if page, err := c.GetInt("page"); err == nil {
		if rows, err := c.GetInt("rows"); err == nil {
			start := (page - 1) * rows
			num := rows
			sql += fmt.Sprintf(" limit %v,%v", start, num)
		}
	}
	//	fmt.Println(sql)
	//	fmt.Println(count_sql)
	_, _ = models.O.Raw(sql).QueryRows(&res)
	var domains models.DomiansResponse
	_ = models.O.Raw(count_sql).QueryRow(&domains.Total)
	domains.Items = res
	c.Data["json"] = domains
	c.ServeJSON(true)
}

type DomainController struct {
	beego.Controller
}

func (c *DomainController) Get() {
	var res models.Records
	var resp Dto
	name := c.Ctx.Input.Param(":splat")
	if !strings.Contains(name, ".") {
		resp.Code = "error"
		resp.Msg = "not a valid domain"
		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}
	res.Name = name

	err := models.O.Read(&res)
	if err == nil {
		c.Data["json"] = res
		c.ServeJSON(true)
	} else {
		resp.Code = "error"
		resp.Msg = "domain not exist"
		c.Data["json"] = resp
		c.ServeJSON(true)
	}
}

//添加域名
func (c *DomainController) Put() {
	name := c.Ctx.Input.Param(":splat")
	content := c.GetString("content")
	ttl, err := c.GetInt("ttl")
	auth := c.GetString("auth")
	var resp Dto
	if !strings.Contains(name, ".") {
		resp.Code = "error"
		resp.Msg = "not a valid domain"
		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}
	if content == "" {
		resp.Code = "error"
		resp.Msg = "content is miss,need ip or cname"
		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}
	if auth == "" {
		resp.Code = "error"
		resp.Msg = "auth is miss,need author name"
		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	if err != nil {
		ttl = 0
	}
	rec := models.Records{}
	rec.Name = name
	rec.Content = content
	rec.Ttl = ttl
	rec.Modifier_ip = c.Ctx.Request.RemoteAddr
	rec.Auth = auth

	err = models.Save(rec)
	if err != nil {
		resp.Code = "error"
		resp.Msg = err.Error()
	} else {
		resp.Code = "ok"
		resp.Msg = "add domain success"
	}
	c.Data["json"] = resp
	c.ServeJSON(true)
}

//更新域名
func (c *DomainController) Post() {
	name := c.Ctx.Input.Param(":splat")
	content := c.GetString("content")
	ttl, err := c.GetInt("ttl")
	auth := c.GetString("auth")
	if err != nil {
		ttl = 0
	}
	var resp Dto
	if content == "" {
		resp.Code = "error"
		resp.Msg = "content is miss,need ip or cname"
		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}
	if auth == "" {
		resp.Code = "error"
		resp.Msg = "auth is miss,need author name"
		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	rec := models.Records{}
	rec.Name = name
	rec.Content = content
	rec.Ttl = ttl
	rec.Modifier_ip = c.Ctx.Request.RemoteAddr
	rec.Auth = auth

	err = models.Save(rec)
	if err == nil {
		resp.Code = "ok"
		resp.Msg = "update domain success"
		c.Data["json"] = resp
		c.ServeJSON(true)
	} else {
		resp.Code = "error"
		resp.Msg = fmt.Sprintf("fail to save info in database with err:%v", err.Error())
		c.Data["json"] = resp
		c.ServeJSON(true)
	}
}

//删除域名
func (c *DomainController) Delete() {
	var rec models.Records
	rec.Name = c.Ctx.Input.Param(":splat")
	var resp Dto

	err := models.Delete(rec)
	if err == nil {
		//			beego.BeeLogger.Info("delete domain %v success", rec.Name)
		resp.Code = "ok"
		resp.Msg = "delete domain success"
		c.Data["json"] = resp
		c.ServeJSON(true)
	} else {
		//			beego.BeeLogger.Info("delete domain %v fail", rec.Name)
		resp.Code = "error"
		resp.Msg = fmt.Sprintf("delete domain fail in mysql operation,with err:%v", err.Error())
		c.Data["json"] = resp
		c.ServeJSON(true)
	}

}
