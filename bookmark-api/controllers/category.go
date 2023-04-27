package controllers

import (
	"encoding/json"
	"github.com/YuZongYangHi/bookmark-platform/bookmark-api/models"
	"strconv"
)

type Category struct {
	BaseController
}

// @router / [get]
func (c *Category) List() {
	result := models.CategoryModel.List()
	c.SuccessResponse(result)
}

// @router /:id [delete]
func (c *Category) Delete() {
	n := c.Ctx.Input.Param(":id")

	pk, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		c.BadRequestResponse("invalid category id")
		return
	}
	_, err = models.CategoryModel.Delete(pk)
	if err != nil {
		c.ErrorResponse(500, err.Error())
		return
	}

	if _, err = models.ItemModel.DeleteByCid(pk); err != nil {
		c.ErrorResponse(500, err.Error())
		return
	}
	
	c.SuccessResponse(pk)
}

// @router /:id [put]
func (c *Category) Put() {
	n := c.Ctx.Input.Param(":id")

	pk, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		c.BadRequestResponse("invalid category id")
		return
	}

	var md models.Category

	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &md); err != nil || md.Name == "" {
		c.BadRequestResponse("name is not null")
		return
	}

	_, err = models.CategoryModel.Update(pk, &md)
	if err != nil {
		c.ErrorResponse(500, err.Error())
		return
	}
	c.SuccessResponse(md)
}

// @router / [post]
func (c *Category) Post() {
	var md models.Category

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &md); err != nil || md.Name == "" {
		c.BadRequestResponse("name is not null")
		return
	}

	if _, err := models.CategoryModel.Create(&md); err != nil {
		c.ErrorResponse(500, err.Error())
		return
	}
	c.SuccessResponse(md)
}
