package dao

import (
	"context"
	"errors"
	"gin_mall_tmp/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

// ListOrderByCondition   用id提取Order
func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Model(&model.Order{}).Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Order("created_at desc").Find(&orders).Error
	return
}

func (dao *OrderDao) DeleteOrderById(uId, oId uint) (err error) {
	//todo:gorm中的删除仅是更新delete_time,若在实际生产中需要再开一个监视应用，不时清理数据库中的数据
	var claimAddress model.Order
	err = dao.DB.Model(&model.Order{}).Where("id=?", oId).First(&claimAddress).Error
	if err != nil {
		return err
	}
	if claimAddress.UserID != uId {
		//todo:按理说这边验证登录应该靠验证token匹配或者交给前端逻辑来做，但是我不太会（ 所以填一个较为粗糙的验证
		return errors.New("用户id不匹配")
	}
	err = dao.DB.Model(&model.Order{}).Where("id=？ AND user_id=?", oId, uId).Delete(&model.Order{}).Error
	return
}

func (dao *OrderDao) CreateOrder(Order *model.Order) (err error) {
	err = dao.DB.Model(&model.Order{}).Create(&Order).Error
	return
}

func (dao *OrderDao) UpdateOrderById(uId uint, oId uint, order *model.Order) error {
	var claimAddress model.Order
	err := dao.DB.Model(&model.Order{}).Where("id=?", oId).First(&claimAddress).Error
	if err != nil {
		return err
	}
	if claimAddress.UserID != uId {
		//todo:按理说这边验证登录应该靠验证token匹配或者交给前端逻辑来做，但是我不太会（ 所以填一个较为粗糙的验证
		return errors.New("用户id不匹配")
	}
	return dao.DB.Model(&model.Order{}).Where("id=?", oId).Updates(&order).Error
}

func (dao *OrderDao) GetOrderById(oId uint, uId uint) (order *model.Order, err error) {
	var claimAddress model.Order
	err = dao.DB.Model(&model.Order{}).Where("id=?", oId).First(&claimAddress).Error
	if err != nil {
		return
	}
	if claimAddress.UserID != uId {
		//todo:按理说这边验证登录应该靠验证token匹配或者交给前端逻辑来做，但是我不太会（ 所以填一个较为粗糙的验证
		return nil, errors.New("用户id不匹配")
	}
	err = dao.DB.Model(&model.Order{}).Where("id=?", oId).First(&order).Error
	return
}

func (dao *OrderDao) OrderExistOrNot(oId uint, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", oId, uId).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
