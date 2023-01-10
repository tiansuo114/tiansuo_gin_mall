package routes

import (
	"gin_mall_tmp/api/v1"
	"gin_mall_tmp/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		v1.GET("user/valid-email/:token", api.ValidEmail)

		// 轮播图
		v1.GET("carousels", api.ListCarousel)

		// 商品操作
		v1.GET("products", api.ListProduct)
		v1.GET("products/:id", api.ShowProduct)
		v1.GET("imgs/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategory)

		authed := v1.Group("/") //group内操作需要登录保护
		authed.Use(middleware.JWT())
		{
			//用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)

			// 显示金额
			authed.POST("money", api.ShowMoney)

			//商品操作
			authed.POST("product", api.CreateProduct)
			authed.PUT("product/:id", api.UpdateProduct)
			authed.DELETE("product/:id", api.DeleteProduct)
			authed.POST("products", api.SearchProductProduct)

			//收藏夹操作
			authed.GET("favorites", api.ListFavorites)
			authed.POST("favorites", api.CreateFavorites)
			authed.DELETE("favorites/:id", api.DeleteFavorites)

			//地址操作
			authed.POST("address", api.CreateAddress)
			authed.GET("addresses/:id", api.GetAddress)
			authed.PUT("addresses/:id", api.UpdateAddress)
			authed.DELETE("addresses/:id", api.DeleteAddress)
			authed.GET("addresses", api.ListAddresses)

			//购物车操作
			authed.POST("carts", api.CreateCart)
			authed.GET("carts/:id", api.GetCart)
			authed.PUT("carts/:id", api.UpdateCart)
			authed.DELETE("carts/:id", api.DeleteCart)
			authed.GET("carts", api.ListCarts)

			//订单操作
			authed.POST("order", api.CreateOrder)
			authed.GET("orders/:id", api.GetOrder)
			authed.DELETE("orders/:id", api.DeleteOrder)
			authed.GET("orders", api.ListOrders)

			//支付功能
			authed.POST("paydown", api.OrderPay)

		}
	}

	return r
}
