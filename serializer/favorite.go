package serializer

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/util"
)

type Favorite struct {
	UserId     uint    `json:"user_id"`
	FavoriteId uint    `json:"favorite_id"`
	CreatedAt  int64   `json:"created_at"`
	Product    Product `json:"product"`
}

func BuildFavorite(favorite *model.Favorite, product *model.Product) Favorite {
	return Favorite{
		UserId:     favorite.UserID,
		FavoriteId: favorite.ID,
		Product:    BuildProduct(product),
		CreatedAt:  favorite.CreatedAt.Unix(),
	}
}

func BuildFavorites(ctx context.Context, items []model.Favorite) (Favorites []Favorite) {
	productDao := dao.NewProductDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			util.LogrusObj.Infoln("get product api", err)
			continue
		}
		Favorite := BuildFavorite(&item, product)
		Favorites = append(Favorites, Favorite)
	}
	return
}
