package controllers

import (
	"encoding/json"
	"github.com/YuZongYangHi/bookmark-platform/bookmark-api/models"
	"strconv"
)

type Item struct {
	BaseController
}

// @router / [get]
func (c *Item) List() {
	result := models.ItemModel.List()
	c.SuccessResponse(result)
}

// @router /:id [get]
func (c *Item) Get() {
	n := c.Ctx.Input.Param(":id")

	pk, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		c.BadRequestResponse("invalid item id")
		return
	}

	result, err := models.ItemModel.Get(pk)
	if err != nil {
		c.ErrorResponse(500, err.Error())
		return
	}
	c.SuccessResponse(result)
}

// @router /:id [delete]
func (c *Item) Delete() {
	n := c.Ctx.Input.Param(":id")

	pk, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		c.BadRequestResponse("invalid item id")
		return
	}
	_, err = models.ItemModel.Delete(pk)
	if err != nil {
		c.ErrorResponse(500, err.Error())
		return
	}
	c.SuccessResponse(pk)
}

// @router /:id [put]
func (c *Item) Put() {
	n := c.Ctx.Input.Param(":id")

	pk, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		c.BadRequestResponse("invalid item id")
		return
	}

	var md models.Item

	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &md); err != nil || md.Name == "" {
		c.BadRequestResponse("name is not null")
		return
	}

	_, err = models.ItemModel.Update(pk, &md)
	if err != nil {
		c.ErrorResponse(500, err.Error())
		return
	}
	c.SuccessResponse(md)
}

// @router / [post]
func (c *Item) Post() {
	var md models.Item

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &md); err != nil {
		c.BadRequestResponse("name is not null")
		return
	}

	if md.IframeURL == "" || md.Name == "" || md.CategoryId == 0 {
		c.BadRequestResponse("invalid params")
		return
	}

	categoryObj, err := models.CategoryModel.Get(md.CategoryId)
	if err != nil || categoryObj.Id == 0 {
		c.BadRequestResponse("category not found")
		return
	}

	if _, err := models.ItemModel.Create(&md); err != nil {
		c.ErrorResponse(500, err.Error())
		return
	}
	c.SuccessResponse(md)
}
