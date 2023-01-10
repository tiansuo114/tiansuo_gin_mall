package dao

import (
	"context"
	"errors"
	"gin_mall_tmp/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

// ListFavorite   用id提取Favorite
func (dao *FavoriteDao) ListFavorite(uid uint) (favorite []model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id=?", &uid).Find(&favorite).Error
	return
}

func (dao *FavoriteDao) DeleteFavorite(uId, fId uint) (err error) {
	//todo:gorm中的删除仅是更新delete_time,若在实际生产中需要再开一个监视应用，不时清理数据库中的数据
	var claimAddress model.Favorite
	err = dao.DB.Model(&model.Favorite{}).Where("id=?", fId).First(&claimAddress).Error
	if err != nil {
		return err
	}
	if claimAddress.UserID != uId {
		//todo:按理说这边验证登录应该靠验证token匹配或者交给前端逻辑来做，但是我不太会（ 所以填一个较为粗糙的验证
		return errors.New("用户id不匹配")
	}
	err = dao.DB.Model(&model.Favorite{}).Where("id=？ AND user_id=?", fId, uId).Delete(&model.Favorite{}).Error
	return
}

func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) (err error) {
	err = dao.DB.Model(&model.Favorite{}).Create(&favorite).Error
	return
}

func (dao *FavoriteDao) FavoriteExistOrNot(pId uint, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id=? AND user_id=?", pId, uId).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
