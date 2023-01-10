package serializer

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/util"
)

type Cart struct {
	ID       uint    `json:"ID"`
	CreateAt int64   `json:"create_at"`
	Product  Product `json:"product"`
	Num      uint    `json:"num"`
	MaxNum   uint    `json:"max_num"`
}

func BuildCart(cart *model.Cart, product *model.Product) Cart {
	return Cart{
		ID:       cart.ID,
		CreateAt: cart.CreatedAt.Unix(),
		Product:  BuildProduct(product),
		Num:      cart.Num,
		MaxNum:   cart.MaxNum,
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) (carts []Cart) {
	productDao := dao.NewProductDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			util.LogrusObj.Infoln("get product api", err)
			continue
		}
		cart := BuildCart(item, product)
		carts = append(carts, cart)
	}
	return
}
