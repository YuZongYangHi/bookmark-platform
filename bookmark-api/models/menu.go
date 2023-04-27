package models

import "fmt"

type Menu struct {
	MenuId int64          `json:"menuId"`
	Path   string         `json:"path"`
	Name   string         `json:"name"`
	Routes []MenuChildren `json:"routes"`
}

type MenuChildren struct {
	IframeURL string `json:"iframeURL"`
	TargetId  int64  `json:"targetId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
}

type menuModel struct{}

func (c *menuModel) List() []*Menu {
	var result []*Menu
	categoryList := CategoryModel.List()
	if len(categoryList) == 0 {
		return result
	}

	for _, category := range categoryList {
		m := &Menu{
			MenuId: category.Id,
			Path:   fmt.Sprintf("/%d", category.Id),
			Name:   category.Name,
			Routes: []MenuChildren{},
		}

		itemList, _ := ItemModel.ListByCategoryId(category.Id)
		if len(itemList) > 0 {
			for _, item := range itemList {
				m.Routes = append(m.Routes, MenuChildren{
					IframeURL: item.IframeURL,
					TargetId:  item.Id,
					Path:      fmt.Sprintf("/%d/%d", category.Id, item.Id),
					Name:      item.Name,
				})
			}
		}
		result = append(result, m)
	}
	return result
}
