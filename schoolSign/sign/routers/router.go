package routers

import (
	"github.com/astaxie/beego"
	"sign/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sign", &controllers.SignController{})
}
