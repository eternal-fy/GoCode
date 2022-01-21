package controllers

import "github.com/astaxie/beego"

type RestfulController struct {
	beego.Controller
}

func (this *RestfulController) PostM() {
	this.Ctx.WriteString("post----")
}

func (this *RestfulController) GetM() {
	this.Ctx.WriteString("get----")
}
