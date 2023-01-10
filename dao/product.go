package dao

import (
	"context"
	"errors"
	"gin_mall_tmp/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// GetProductById 用id获得Product信息
func (dao *ProductDao) GetProductById(id uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id=?", id).First(&product).Error
	return
}

func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Where(condition).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&products).Error
	return
}

func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, count int64, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Count(&count).Error
	if err != nil {
		return
	}

	err = dao.DB.Model(&model.Product{}).
		Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * (page.PageSize)).
		Limit(page.PageSize).Find(&products).Error
	return
}

func (dao *ProductDao) DeleteProductById(pId uint, uId uint) error {
	//todo:gorm中的删除仅是更新delete_time,若在实际生产中需要再开一个监视应用，不时清理数据库中的数据
	var claimAddress model.Product
	err := dao.DB.Model(&model.Product{}).Where("id=?", pId).First(&claimAddress).Error
	if err != nil {
		return err
	}
	if claimAddress.BossID != uId {
		//todo:按理说这边验证登录应该靠验证token匹配或者交给前端逻辑来做，但是我不太会（ 所以填一个较为粗糙的验证
		return errors.New("用户id不匹配")
	}
	err = dao.DB.Model(&model.Product{}).Where("id=?", pId).Delete(&model.Product{}).Error
	return err
}

func (dao *ProductDao) UpdateProductById(uId uint, pId uint, product *model.Product) error {
	var claimAddress model.Product
	err := dao.DB.Model(&model.Product{}).Where("id=?", pId).First(&claimAddress).Error
	if err != nil {
		return err
	}
	if claimAddress.BossID != uId {
		//todo:按理说这边验证登录应该靠验证token匹配或者交给前端逻辑来做，但是我不太会（ 所以填一个较为粗糙的验证
		return errors.New("用户id不匹配")
	}
	return dao.DB.Model(&model.Product{}).Where("id=?", pId).Updates(&product).Error
}

func (dao *ProductDao) ProductExistOrNot(pId uint, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Product{}).Where("product_id=? AND user_id=?", pId, uId).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (dao *ProductDao) GetProductByName(name string, uId uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("name=? AND user_id=?", name, uId).First(&product).Error
	return
}
