package main

import (
	_ "dnsadmin/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SessionOn = true
	beego.Run()
}
