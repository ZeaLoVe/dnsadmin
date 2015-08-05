package controllers

import (
	"dnsadmin/models"
	"github.com/astaxie/beego"
)

type SearchController struct {
	beego.Controller
}

func (c *SearchController) Post() {
	c.Data["Website"] = "DNSadmin"

	var res []models.Records
	target := c.GetString("searchDNS")
	//	c.Ctx.WriteString(target)
	if target == "" {
		_, _ = models.O.Raw("select * from records").QueryRows(&res)
	} else {
		res, _ = models.Search(target)
	}
	c.Data["s"] = res
	c.Layout = "index.tpl"
	c.TplNames = "data.tpl"
}
