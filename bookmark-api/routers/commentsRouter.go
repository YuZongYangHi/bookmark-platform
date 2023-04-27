package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Category"] = append(beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Category"],
        beego.ControllerComments{
            Method: "List",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Category"] = append(beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Category"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Category"] = append(beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Category"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Category"] = append(beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Category"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Item"] = append(beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Item"],
        beego.ControllerComments{
            Method: "List",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Item"] = append(beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Item"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Item"] = append(beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Item"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Item"] = append(beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Item"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Menu"] = append(beego.GlobalControllerRouter["github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers:Menu"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
