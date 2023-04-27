package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

const CategoryTableName = "category"

type Category struct {
	Id      int64     `orm:"pk;auto; column(id)" json:"id"`
	Name    string    `orm:"column(name);size(255)" json:"name"`
	Created time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
	Updated time.Time `orm:"auto_now;type(datetime)" json:"updated"`
}

type categoryModel struct{}

func (u *Category) TableName() string {
	return CategoryTableName
}

func (c *categoryModel) List() []*Category {
	var result []*Category
	Orm().QueryTable(CategoryTableName).All(&result)
	return result
}

func (c *categoryModel) Create(md *Category) (int64, error) {
	return Orm().Insert(md)
}

func (c *categoryModel) Update(pk int64, md *Category) (int64, error) {
	v := orm.Params{
		"name":    md.Name,
		"updated": time.Now(),
	}
	return Orm().QueryTable(CategoryTableName).Filter("id", pk).Update(v)
}

func (c *categoryModel) Delete(pk int64) (int64, error) {
	return Orm().QueryTable(CategoryTableName).Filter("id", pk).Delete()
}

func (c *categoryModel) Get(pk int64) (*Category, error) {
	var result Category
	return &result, Orm().QueryTable(CategoryTableName).Filter("id", pk).One(&result)
}
