package controllers

import (
	"github.com/astaxie/beego"
	"sign/service"
	"strings"
)

type MainController struct {
	beego.Controller
}
type SignController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *SignController) Get() {
	fetch := service.Fetch("http://180.76.172.177/18nodone")
	split := strings.Split(fetch, "<br>")
	for _, context := range split {
		strings.Split(context, "' '+")
	}
	c.Ctx.WriteString(fetch)
}
