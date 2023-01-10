package dao

import (
	"context"
	"errors"
	"gin_mall_tmp/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

// ListCart   用id提取Cart
func (dao *CartDao) ListCart(uid uint) (Cart []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id=?", &uid).Find(&Cart).Error
	return
}

func (dao *CartDao) DeleteCart(uId, cId uint) (err error) {
	//todo:gorm中的删除仅是更新delete_time,若在实际生产中需要再开一个监视应用，不时清理数据库中的数据
	var claimAddress model.Cart
	err = dao.DB.Model(&model.Cart{}).Where("id=?", cId).First(&claimAddress).Error
	if err != nil {
		return err
	}
	if claimAddress.UserID != uId {
		//todo:按理说这边验证登录应该靠验证token匹配或者交给前端逻辑来做，但是我不太会（ 所以填一个较为粗糙的验证
		return errors.New("用户id不匹配")
	}
	err = dao.DB.Model(&model.Cart{}).Where("id=？ AND user_id=?", cId, uId).Delete(&model.Cart{}).Error
	return
}

func (dao *CartDao) CreateCart(Cart *model.Cart) (err error) {
	err = dao.DB.Model(&model.Cart{}).Create(&Cart).Error
	return
}

func (dao *CartDao) CartExistOrNot(pId uint, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Cart{}).Where("cart_id=? AND user_id=?", pId, uId).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (dao *CartDao) GetCartById(id uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=?", id).First(&cart).Error
	return
}

func (dao *CartDao) UpdateCartById(uId uint, cId uint, cart *model.Cart) error {
	var claimAddress model.Cart
	err := dao.DB.Model(&model.Cart{}).Where("id=?", cId).First(&claimAddress).Error
	if err != nil {
		return err
	}
	if claimAddress.UserID != uId {
		//todo:按理说这边验证登录应该靠验证token匹配或者交给前端逻辑来做，但是我不太会（ 所以填一个较为粗糙的验证
		return errors.New("用户id不匹配")
	}
	return dao.DB.Model(&model.Cart{}).Where("id=?", cId).Updates(&cart).Error
}
