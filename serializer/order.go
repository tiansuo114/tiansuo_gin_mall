package serializer

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/util"
)

type Order struct {
	ID        uint    `json:"id"`
	OrderNum  uint64  `json:"order_num"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Num       uint    `json:"num"`
	Address   Address `json:"address"`
	Type      uint    `json:"type"`
	Product   Product `json:"product"`
}

func BuildOrder(item1 *model.Order, item2 *model.Product, item3 *model.Address) Order {
	return Order{
		ID:        item1.ID,
		OrderNum:  item1.OrderNum,
		CreatedAt: item1.CreatedAt.Unix(),
		UpdatedAt: item1.UpdatedAt.Unix(),
		UserID:    item1.UserID,
		ProductID: item1.ProductID,
		Num:       uint(item1.Num),
		Address:   BuildAddress(item3),
		Type:      item1.Type,
		Product:   BuildProduct(item2),
	}
}

func BuildOrders(ctx context.Context, items []*model.Order) (orders []Order) {
	productDao := dao.NewProductDao(ctx)
	addressDao := dao.NewAddressDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			util.LogrusObj.Infoln("build order api", err)
		}
		address, err := addressDao.GetAddressById(item.AddressID)
		if err != nil {
			util.LogrusObj.Infoln("build order api", err)
		}
		order := BuildOrder(item, product, address)
		orders = append(orders, order)
	}
	return
}
