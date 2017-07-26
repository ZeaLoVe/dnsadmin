package controllers

import (
	"dnsadmin/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	sess := c.StartSession()
	username := sess.Get("user_name")
	if username == "" {
		c.Ctx.Redirect(302, "/")
		return
	}
	c.Data["UserName"] = username
	c.Data["Website"] = "DNSadmin"
	c.TplName = "index.tpl"
}

func (c *MainController) redirect_to_sso() {
	sig, err := c.genSig()
	c.Ctx.SetCookie("sig", sig)
	//fmt.Printf("sig from gensig:%v\n", sig)
	if err != nil {
		return
	}
	//	callback := fmt.Sprintf("http://%v", beego.AppConfig.String("domain"), beego.AppConfig.String("httpport"))
	callback := fmt.Sprintf("http://%v", c.Ctx.Request.Host)
	c.Redirect(c.getLoginUrl(sig, callback), 302)
}

func (c *MainController) Prepare() {
	sess := c.StartSession()
	flag, _ := beego.AppConfig.Bool("auth")
	if flag == false {
		c.SetSession("user_name", "NotNeedAuth")
	}
	user_name := sess.Get("user_name")
	if user_name == nil {
		sig := c.Ctx.GetCookie("sig")
		if sig == "" {
			c.redirect_to_sso()
			return
		}
		username := c.username_form_sso(sig)
		//fmt.Println(username)
		if username == "" {
			c.redirect_to_sso()
			return
		}
		c.SetSession("user_name", username)
		return
	} else {
		//already login before
	}
}

func (c *MainController) getLoginUrl(sig string, callback string) string {
	return fmt.Sprintf("%v:%v/auth/login?sig=%v&callback=%v", beego.AppConfig.String("uic"), beego.AppConfig.String("uicport"), sig, callback)
}

func (c *MainController) genSig() (string, error) {
	url := fmt.Sprintf("%v:%v/sso/sig", beego.AppConfig.String("uic"), beego.AppConfig.String("uicport"))
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (c *MainController) username_form_sso(sig string) string {
	url := fmt.Sprintf("%v:%v/sso/user/%v?token=%v", beego.AppConfig.String("uic"), beego.AppConfig.String("uicport"), sig, beego.AppConfig.String("uictoken"))
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return ""
	}
	if body != nil {
		defer resp.Body.Close()
	}
	var userinfoResp models.UersInfoResponse
	err = json.Unmarshal(body, &userinfoResp)
	if userinfoResp.UserInfo.Name == "" {
		return ""
	}
	return userinfoResp.UserInfo.Name
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Ctx.WriteString(fmt.Sprint(c.GetStrings("Counter")))
	c.Ctx.WriteString(fmt.Sprint(c.GetString("Endpoint")))
}
