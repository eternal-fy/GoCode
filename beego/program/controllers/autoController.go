package controllers

import "github.com/astaxie/beego"

type AutoController struct {
	beego.Controller
}

func (this *AutoController) Login() {
	this.Ctx.WriteString("login-----")
}

func (this *AutoController) Logout() {
	this.Ctx.WriteString("logout-----")
}
