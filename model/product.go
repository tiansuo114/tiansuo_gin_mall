package model

import (
	"fmt"
	//"mall/cache"
	//"strconv"
	"gin_mall_tmp/cache"
	"gorm.io/gorm"
	"strconv"
)

//商品模型
type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryID    uint   `gorm:"not null"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

//View 获取点击数
func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	fmt.Println(cache.ProductViewKey(product.ID))
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 增加商品点击数
func (product *Product) AddView() {
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
