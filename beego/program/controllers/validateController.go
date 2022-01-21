package controllers

import "github.com/astaxie/beego"

type ValidateController struct {
	beego.Controller
}

func (this *ValidateController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	if username == "ld" && password == "ld" {
		this.SetSession("ld", 100)
		this.Ctx.Redirect(302, "/mystatic/welcome.html")
	}
}
