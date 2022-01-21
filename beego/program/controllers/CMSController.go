package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type CMSController struct {
	beego.Controller
}

// @router /r1/:key [get]
func (this *CMSController) StaticBlock() {
	getString := this.GetString("name")
	this.Ctx.WriteString(getString)
}

type user struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}

// @router /user/:key [get]
func (this *CMSController) AllBlock() {
	this.Ctx.Input.Param(":key")
	u := user{}
	this.ParseForm(&u)
	this.Ctx.WriteString(fmt.Sprintf("%v", u))
	//传递json数据
	mystruct := user{1, "ld", 10, ""}
	this.Data["json"] = &mystruct
	this.ServeJSON()
}
