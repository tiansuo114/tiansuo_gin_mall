package middleware

import (
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var err error
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ErrorJwtMiddleWare
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.Success {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"error":  err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
