package controllers

import "github.com/YuZongYangHi/bookmark-platform/bookmark-api/models"

type Menu struct {
	BaseController
}

// @router /list [get]
func (c *Menu) List() {
	result := models.MenuModel.List()
	c.SuccessResponse(result)
}
