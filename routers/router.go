package routers

import (
	"dnsadmin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/disable/*", &controllers.DisableController{})
	beego.Router("/enable/*", &controllers.EnableController{})
	beego.Router("/syncall", &controllers.SyncAllController{})
	beego.Router("/domains", &controllers.DomainsController{})
	beego.Router("/domains/*", &controllers.DomainController{})
	beego.Router("/test", &controllers.TestController{})
}
