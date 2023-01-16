package api

import (
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFavorites(c *gin.Context) {
	createFavoritesService := service.FavoriteService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createFavoritesService); err == nil {
		res := createFavoritesService.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Favorites api", err)
	}
}

func ListFavorites(c *gin.Context) {
	listFavoritesService := service.FavoriteService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listFavoritesService); err == nil {
		res := listFavoritesService.List(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("list Favorites api", err)
	}
}

func DeleteFavorites(c *gin.Context) {
	deleteFavoriteService := service.FavoriteService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteFavoriteService); err == nil {
		res := deleteFavoriteService.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("delete favorite api", err)
	}

}
