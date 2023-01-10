package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `form:"id" json:"id"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	ProductId uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
}

func (service *CartService) List(ctx context.Context, uId uint) serializer.Response {
	CartDao := dao.NewCartDao(ctx)
	code := e.Success
	Carts, err := CartDao.ListCart(uId)
	if err != nil {
		util.LogrusObj.Infoln("list Carts api", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarts(ctx, Carts), uint(len(Carts)))
}
func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	cartDao := dao.NewCartDao(ctx)
	code := e.Success
	exist, err := cartDao.CartExistOrNot(service.ProductId, uId)
	if exist {
		code = e.ErrorCartExist
		util.LogrusObj.Infoln("ErrorCartExist", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("Cart find product api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cart := &model.Cart{
		UserID:    uId,
		ProductID: product.ID,
		BossID:    product.BossID,
		Num:       1,
		MaxNum:    uint(product.Num),
		Check:     false,
	}
	err = cartDao.CreateCart(cart)
	if err != nil {
		code = e.ErrorCartCreate
		util.LogrusObj.Infoln("Cart create api", err)
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
func (service *CartService) Delete(ctx context.Context, uId uint, cId string) serializer.Response {
	cartDao := dao.NewCartDao(ctx)
	code := e.Success
	cartId, _ := strconv.Atoi(cId)
	err := cartDao.DeleteCart(uId, uint(cartId))
	if err != nil {
		code = e.ErrorCartDelete
		util.LogrusObj.Infoln("cart delete api", err)
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
func (service *CartService) UpDate(ctx context.Context, uId uint, cId string) serializer.Response {
	cartDao := dao.NewCartDao(ctx)
	code := e.Success
	cartId, _ := strconv.Atoi(cId)
	cart, err := cartDao.GetCartById(uint(cartId))
	if err != nil {
		code = e.ErrorCartGet
		util.LogrusObj.Infoln("cart update api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(cart.ProductID)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("cart update api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cart.Num = service.Num
	err = cartDao.UpdateCartById(uId, uint(cartId), cart)
	if err != nil {
		code = e.ErrorCartUpdate
		util.LogrusObj.Infoln("cart update api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product),
	}
}
func (service *CartService) Get(ctx context.Context, cId string) serializer.Response {
	cartDao := dao.NewCartDao(ctx)
	code := e.Success
	cartId, _ := strconv.Atoi(cId)
	cart, err := cartDao.GetCartById(uint(cartId))
	if err != nil {
		code = e.ErrorCartGet
		util.LogrusObj.Infoln("cart get api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(cart.ProductID)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product),
	}
}
