package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

const ItemTableName = "item"

type Item struct {
	Id          int64     `orm:"pk;auto; column(id)" json:"id"`
	Name        string    `orm:"column(name);size(255)" json:"name"`
	CategoryId  int64     `orm:"column(category_id)" json:"categoryId"`
	IframeURL   string    `orm:"column(iframe_url);type(text)" json:"iframeURL"`
	Description string    `orm:"type(text)" json:"description"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

type itemModel struct{}

func (u *Item) TableName() string {
	return ItemTableName
}

func (c *itemModel) List() []*Item {
	var result []*Item
	Orm().QueryTable(ItemTableName).All(&result)
	return result
}

func (c *itemModel) ListByCategoryId(cid int64) ([]*Item, error) {
	var result []*Item
	_, err := Orm().QueryTable(ItemTableName).Filter("categoryId", cid).All(&result)
	return result, err
}

func (c *itemModel) Create(md *Item) (int64, error) {
	return Orm().Insert(md)
}

func (c *itemModel) Update(pk int64, md *Item) (int64, error) {
	v := orm.Params{
		"name":        md.Name,
		"iframe_url":  md.IframeURL,
		"description": md.Description,
		"updated":     time.Now(),
	}
	return Orm().QueryTable(ItemTableName).Filter("id", pk).Update(v)
}

func (c *itemModel) Delete(pk int64) (int64, error) {
	return Orm().QueryTable(ItemTableName).Filter("id", pk).Delete()
}

func (c *itemModel) Get(pk int64) (*Item, error) {
	var obj Item
	return &obj, Orm().QueryTable(ItemTableName).Filter("id", pk).One(&obj)
}

func (c *itemModel) DeleteByCid(cid int64) (int64, error) {
	return Orm().QueryTable(ItemTableName).Filter("categoryId", cid).Delete()
}
