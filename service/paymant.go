package service

import (
	"context"
	"errors"
	"fmt"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/serializer"
	"strconv"
)

type OrderPay struct {
	OrderId   uint    `form:"order_id" json:"order_id"`
	Money     float64 `form:"money" json:"money"`
	OrderNo   string  `form:"orderNo" json:"orderNo"`
	ProductID int     `form:"product_id" json:"product_id"`
	PayTime   string  `form:"payTime" json:"payTime" `
	Sign      string  `form:"sign" json:"sign" `
	BossID    int     `form:"boss_id" json:"boss_id"`
	BossName  string  `form:"boss_name" json:"boss_name"`
	Num       int     `form:"num" json:"num"`
	Key       string  `form:"key" json:"key"`
}

func (service OrderPay) PayDown(ctx context.Context, uId uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	tx := orderDao.Begin()
	order, err := orderDao.GetOrderById(service.OrderId, uId)
	if err != nil {
		util.LogrusObj.Infoln("pay order api_order", err)
		code = e.ErrorOrderGet
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num)

	//用户
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		util.LogrusObj.Infoln("pay order api_user", err)
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	moneyStr := util.Encrypt.AesEncoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)

	if moneyFloat-money < 0.0 {
		tx.Rollback()
		util.LogrusObj.Infoln("pay order api_user", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}
	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finMoney)

	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		tx.Rollback()
		util.LogrusObj.Infoln("pay order api_user", err)
		code = e.ErrorUpdateUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//商家
	var boss *model.User
	boss, err = userDao.GetUserById(uint(service.BossID))
	if err != nil {
		tx.Rollback()
		util.LogrusObj.Infoln("pay order api_boss", err)
		code = e.ErrorGetUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	moneyStr = util.Encrypt.AesEncoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)
	err = userDao.UpdateUserById(uint(service.BossID), boss)
	if err != nil {
		tx.Rollback()
		util.LogrusObj.Infoln("pay order api_boss", err)
		code = e.ErrorUpdateUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//商品
	var product *model.Product
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(uint(service.ProductID))
	if err != nil {
		tx.Rollback()
		util.LogrusObj.Infoln("pay order api_product", err)
		code = e.ErrorProductGet
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product.Num -= num

	err = productDao.UpdateProductById(uint(service.BossID), uint(service.ProductID), product)
	if err != nil {
		tx.Rollback()
		util.LogrusObj.Infoln("pay order api_product", err)
		code = e.ErrorProductUpDate
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = orderDao.DeleteOrderById(uId, service.OrderId)
	if err != nil {
		tx.Rollback()
		util.LogrusObj.Infoln("pay order api_order", err)
		code = e.ErrorProductDelete
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//自己的商品+1
	//todo:正常应该不是这样，但是偷懒而且不知道怎么做
	exists, err := productDao.ProductExistOrNot(product.ID, uId)
	if err != nil {
		tx.Rollback()
		util.LogrusObj.Infoln("pay order api_order exists", err)
		code = e.ErrorProductGet
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if exists {
		userProduct, err := productDao.GetProductByName(product.Name, uId)
		if err != nil {
			tx.Rollback()
			util.LogrusObj.Infoln("pay order api exists", err)
			code = e.ErrorProductGet
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		userProduct.Num += num
		err = productDao.UpdateProductById(uId, userProduct.ID, userProduct)
		if err != nil {
			tx.Rollback()
			util.LogrusObj.Infoln("pay order api exists", err)
			code = e.ErrorProductUpDate
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		userProduct := &model.Product{
			Name:          product.Name,
			CategoryID:    product.CategoryID,
			Title:         product.Title,
			Info:          product.Info,
			ImgPath:       product.ImgPath,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
			OnSale:        false,
			Num:           num,
			BossID:        uId,
			BossName:      user.UserName,
			BossAvatar:    user.Avatar,
		}
		err := productDao.CreateProduct(userProduct)
		if err != nil {
			tx.Rollback()
			util.LogrusObj.Infoln("pay order api exists", err)
			code = e.ErrorProductCreate
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
