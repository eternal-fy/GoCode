package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type SessionController struct {
	beego.Controller
}

// @router /session/:key [get]
func (this *SessionController) GetSessionNum() {
	v := this.GetSession("asta")
	if v == nil {
		this.SetSession("asta", int(1))
		this.Data["num"] = 0
	} else {
		this.SetSession("asta", v.(int)+1)
		this.Data["num"] = v.(int)
	}
	this.Ctx.WriteString(fmt.Sprintf("%v", this.Data["num"]))
}
