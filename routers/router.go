package routers

import (
	"dnsadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/add", &controllers.AddController{})
	beego.Router("/insert", &controllers.InsertController{})
	beego.Router("/search", &controllers.SearchController{})
	beego.Router("/disable/*", &controllers.DisableController{})
	beego.Router("/enable/*", &controllers.EnableController{})
	beego.Router("/syncall", &controllers.SyncAllController{})
}
