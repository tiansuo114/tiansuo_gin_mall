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

type AddressService struct {
	Address string `json:"address" form:"address"`
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
}

// Create 创建地址
func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.CreateAddress(address)
	if err != nil {
		code = e.ErrorAddressCreate
		util.LogrusObj.Infoln("address create api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}
}

// Delete 删除地址
func (service *AddressService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addressId, _ := strconv.Atoi(aId)
	err := addressDao.DeleteAddressById(uint(addressId), uId)
	if err != nil {
		code = e.ErrorAddressDelete
		util.LogrusObj.Infoln("address delete api", err)
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

// Update 修改地址
func (service *AddressService) Update(ctx context.Context, uId uint, aId string) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addressId, _ := strconv.Atoi(aId)
	address := &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.UpdateAddressById(uint(addressId), uId, address)
	if err != nil {
		code = e.ErrorAddressUpdate
		util.LogrusObj.Infoln("address update api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}
}

// Get 获取地址
func (service *AddressService) Get(ctx context.Context, aId string) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addressId, _ := strconv.Atoi(aId)
	address, err := addressDao.GetAddressById(uint(addressId))
	if err != nil {
		code = e.ErrorAddressGet
		util.LogrusObj.Infoln("address get api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}

}

// List 获取地址列表
func (service *AddressService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressesDao := dao.NewAddressDao(ctx)
	addresses, err := addressesDao.ListAddressById(uId)
	if err != nil {
		code = e.ErrorAddressList
		util.LogrusObj.Infoln("address list api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddresses(addresses),
	}
}
