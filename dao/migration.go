package dao

import (
	"fmt"
	"gin_mall_tmp/model"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{},
			&model.Product{},
			&model.Carousel{},
			&model.Category{},
			&model.Favorite{},
			&model.ProductImg{},
			&model.Order{},
			&model.Cart{},
			&model.Admin{},
			&model.Address{},
			&model.Notice{},
			&model.SkillGoods{})
	if err != nil {
		fmt.Println("err", err)
	}
	return
}
