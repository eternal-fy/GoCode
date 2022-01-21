package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func MyFilter() beego.FilterFunc {
	return func(ctx *context.Context) {

		key := ctx.Input.Session("ld")
		println(key)
		if key == nil && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/logout")
		}
	}
}
