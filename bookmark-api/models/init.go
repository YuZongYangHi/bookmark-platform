package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var (
	GlobalOrm     orm.Ormer
	once          sync.Once
	CategoryModel *categoryModel
	ItemModel     *itemModel
	MenuModel     *menuModel
)

func init() {
	fmt.Println(123123123)
	orm.RegisterModel(
		new(Category),
		new(Item),
	)

	CategoryModel = &categoryModel{}
	ItemModel = &itemModel{}
	MenuModel = &menuModel{}
}

func Orm() orm.Ormer {
	once.Do(func() {
		GlobalOrm = orm.NewOrm()
	})
	return GlobalOrm
}
