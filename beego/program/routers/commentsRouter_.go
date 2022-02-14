package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["program/controllers:CMSController"] = append(beego.GlobalControllerRouter["program/controllers:CMSController"],
        beego.ControllerComments{
            Method: "StaticBlock",
            Router: "/r1/:key",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["program/controllers:CMSController"] = append(beego.GlobalControllerRouter["program/controllers:CMSController"],
        beego.ControllerComments{
            Method: "AllBlock",
            Router: "/user/:key",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["program/controllers:SessionController"] = append(beego.GlobalControllerRouter["program/controllers:SessionController"],
        beego.ControllerComments{
            Method: "GetSessionNum",
            Router: "/session/:key",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
