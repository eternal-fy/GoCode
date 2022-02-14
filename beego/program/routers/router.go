package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	. "program/controllers"
)

func init() {
	beego.Router("/", &MainController{})
	beego.Router("/login", &ValidateController{})
	beego.Get("/logout", func(context *context.Context) {
		context.Redirect(302, "/mystatic/login.html")
	})
	beego.InsertFilter("/*", beego.BeforeRouter, MyFilter())
	beego.Router("/restful/get", &RestfulController{}, "get:GetM")
	beego.Router("/restful/post", &RestfulController{}, "post:PostM")
	beego.Include(&CMSController{}, &SessionController{}) //注册注解路由，所有的注册路由controller必须放在一起，如果没有生成可能是beego版本低了，需要升级
	beego.AutoRouter(&AutoController{})                   //http://localhost:8080/auto/logout 可以进入

}
