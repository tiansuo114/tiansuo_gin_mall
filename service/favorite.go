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

type FavoriteService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossID     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"boss_id"`
	model.BasePage
}

func (service *FavoriteService) List(ctx context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	favorites, err := favoriteDao.ListFavorite(uId)
	if err != nil {
		util.LogrusObj.Infoln("list favorites api", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(len(favorites)))
}

func (service *FavoriteService) Create(ctx context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	exist, err := favoriteDao.FavoriteExistOrNot(service.ProductId, uId)
	if exist {
		code = e.ErrorFavoriteExist
		util.LogrusObj.Infoln("ErrorFavoriteExist", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("favorite find user api", err)
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
		util.LogrusObj.Infoln("favorite find product api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(product.BossID)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("favorite find boss api", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//todo:bossid 不应该依靠相应传入，应该在product里查找
	favorite := &model.Favorite{
		User:      *user,
		UserID:    uId,
		Product:   *product,
		ProductID: service.ProductId,
		Boss:      *boss,
		BossID:    boss.ID,
	}
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.ErrorFavoriteCreate
		util.LogrusObj.Infoln("favorite create api", err)
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
func (service *FavoriteService) Delete(ctx context.Context, uId uint, fId string) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	favoriteId, _ := strconv.Atoi(fId)
	err := favoriteDao.DeleteFavorite(uId, uint(favoriteId))
	if err != nil {
		code = e.ErrorFavoriteDelete
		util.LogrusObj.Infoln("favorite delete api", err)
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
