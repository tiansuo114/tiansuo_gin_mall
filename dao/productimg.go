package dao

import (
	"context"
	"gin_mall_tmp/model"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

// GetProductImgById 用id获得ProductImg信息
func (dao *ProductImgDao) GetProductImgById(id uint) (productImg *model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("id=?", id).First(&productImg).Error
	return
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

func (dao *ProductImgDao) ListProductImg(id uint) (productImgs []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id=?", id).Find(&productImgs).Error
	return
}
