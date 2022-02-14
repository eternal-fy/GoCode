package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func MyFilter() beego.FilterFunc {
	return func(ctx *context.Context) {
		ctx.Input.Cookie("sfa")
		key := ctx.Input.Session("ld")
		println(key)
		println(ctx.Request.RequestURI)
		if key == nil && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/mystatic/login.html")
		}
	}
}
