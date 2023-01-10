package dao

import (
	"context"
	"errors"
	"gin_mall_tmp/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

// ListAddressById   用uid提取Address
func (dao *AddressDao) ListAddressById(uid uint) (addresses []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id=?", &uid).Find(&addresses).Error
	return
}

func (dao *AddressDao) DeleteAddressById(aId uint, uId uint) (err error) {
	//todo:gorm中的删除仅是更新delete_time,若在实际生产中需要再开一个监视应用，不时清理数据库中的数据
	var claimAddress model.Address
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).First(&claimAddress).Error
	if claimAddress.UserID != uId {
		//todo:按理说这边验证登录应该靠验证token匹配或者交给前端逻辑来做，但是我不太会（ 所以填一个较为粗糙的验证
		return errors.New("用户id不匹配")
	}
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).Delete(&model.Address{}).Error
	return
}

func (dao *AddressDao) CreateAddress(address *model.Address) (err error) {
	err = dao.DB.Model(&model.Address{}).Create(&address).Error
	return
}

func (dao *AddressDao) UpdateAddressById(aId uint, uId uint, address *model.Address) (err error) {
	var claimAddress model.Address
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).First(&claimAddress).Error
	if claimAddress.UserID != uId {
		//todo:按理说这边验证登录应该靠验证token匹配或者交给前端逻辑来做，但是我不太会（ 所以填一个较为粗糙的验证
		return errors.New("用户id不匹配")
	}
	return dao.DB.Model(&model.Address{}).Where("id=?", aId).Updates(&address).Error
}

func (dao *AddressDao) GetAddressById(aId uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).First(&address).Error
	return
}
