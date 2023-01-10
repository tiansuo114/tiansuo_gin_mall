package service

import (
	"context"
	"fmt"
	"gin_mall_tmp/cache"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/serializer"
	"github.com/go-redis/redis"
	"math/rand"
	"strconv"
	"time"
)

const OrderTimeKey = "OrderTime"

type OrderService struct {
	ProductID uint `json:"product_id" form:"product_id"`
	Num       uint `form:"num" json:"num"`
	AddressID uint `form:"address_id" json:"address_id"`
	Money     int  `form:"money" json:"money"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	UserID    uint `form:"user_id" json:"user_id"`
	OrderNum  uint `form:"order_num" json:"order_num"`
	Type      int  `form:"type" json:"type"`
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	order := &model.Order{
		UserID:    uId,
		ProductID: service.ProductID,
		BossID:    service.BossID,
		Num:       int(service.Num),
		Money:     float64(service.Money),
		Type:      1,
	}
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(service.AddressID)
	if err != nil {
		util.LogrusObj.Infoln("create order api", err)
		code = e.ErrorAddressGet
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressID = address.ID
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(service.ProductID))
	userNum := strconv.Itoa(int(uId))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum

	orderDao := dao.NewOrderDao(ctx)
	err = orderDao.CreateOrder(order)
	if err != nil {
		util.LogrusObj.Infoln("create order api", err)
		code = e.ErrorOrderCreate
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//订单号存入Redis中，设置过期时间
	data := redis.Z{
		Score:  float64(time.Now().Unix()) + 15*time.Minute.Seconds(),
		Member: orderNum,
	}
	cache.RedisClient.ZAdd(OrderTimeKey, data)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
func (service *OrderService) DeleteOrderById(ctx context.Context, uId uint, oId string) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	orderId, _ := strconv.Atoi(oId)
	err := orderDao.DeleteOrderById(uId, uint(orderId))
	if err != nil {
		util.LogrusObj.Infoln("delete order api", err)
		code = e.ErrorOrderDelete
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
func (service *OrderService) ListOrderById(ctx context.Context, uId uint) serializer.Response {
	var orders []*model.Order
	var total int64
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 5
	}

	orderDao := dao.NewOrderDao(ctx)
	condition := make(map[string]interface{})
	condition["user_id"] = uId

	if service.Type == 0 {
		condition["type"] = 0
	} else {
		condition["type"] = service.Type
	}
	orders, total, err := orderDao.ListOrderByCondition(condition, service.BasePage)
	if err != nil {
		code = e.ErrorOrderList
		util.LogrusObj.Infoln("list order api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orders), uint(total))

}
func (service *OrderService) GetOrderById(ctx context.Context, uId uint, oId string) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	orderId, _ := strconv.Atoi(oId)
	order, err := orderDao.GetOrderById(uId, uint(orderId))
	if err != nil {
		util.LogrusObj.Infoln("get order api", err)
		code = e.ErrorOrderGet
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	addressDao := dao.NewAddressDao(ctx)
	productDao := dao.NewProductDao(ctx)
	address, err := addressDao.GetAddressById(order.AddressID)
	if err != nil {
		util.LogrusObj.Infoln("get order api", err)
		code = e.ErrorAddressGet
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product, err := productDao.GetProductById(order.ProductID)
	if err != nil {
		util.LogrusObj.Infoln("get order api", err)
		code = e.ErrorProductGet
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, product, address),
	}
}
