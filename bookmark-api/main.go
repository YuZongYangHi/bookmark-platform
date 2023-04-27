package main

import (
	"github.com/YuZongYangHi/bookmark-platform/bookmark-api/initial"
	_ "github.com/YuZongYangHi/bookmark-platform/bookmark-api/models"
	_ "github.com/YuZongYangHi/bookmark-platform/bookmark-api/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	initial.InitDb()
	beego.Run()
}
