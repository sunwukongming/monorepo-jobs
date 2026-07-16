package middleware

import (
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

// AuthC 端账号验证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		_, err := services.AuthDecodeToken(c)
		if err != nil {
			c.Abort()
			services.ResponseError(c, 401, "Token is invalid: token is "+tokenString, nil)
			return
		}
		c.Next()
	}
}
