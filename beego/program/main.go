package main

import (
	"github.com/astaxie/beego"
	_ "program/routers"
)

/*
常见的奇怪错误，方法名首字母没有大写，所以无法通过反射获取，导致报错。
*/
func main() {
	beego.SetStaticPath("/mystatic", "static/html")
	beego.BConfig.WebConfig.Session.SessionOn = true //或者在app.conf中添加session=true
	beego.Run()
}
