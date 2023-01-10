package serializer

import (
	"gin_mall_tmp/model"
)

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		Id:           item.ID,
		CategoryName: item.CategoryName,
	}
}

func BuildCategories(items []model.Category) (categories []Category) {
	for _, item := range items {
		category := BuildCategory(&item)
		categories = append(categories, category)
	}
	return
}
