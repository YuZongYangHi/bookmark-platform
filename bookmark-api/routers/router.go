package routers

import (
	"github.com/YuZongYangHi/bookmark-platform/bookmark-api/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	apiNamespaceRouters := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/category",
			beego.NSInclude(&controllers.Category{}),
		),
		beego.NSNamespace("/item",
			beego.NSInclude(&controllers.Item{}),
		),
		beego.NSNamespace("/menu",
			beego.NSInclude(&controllers.Menu{}),
		),
	)

	beego.AddNamespace(apiNamespaceRouters)
}
