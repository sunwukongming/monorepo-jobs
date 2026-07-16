package services

import (
	"app/config"
	"app/db"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthDecodeToken(c *gin.Context) (*jwt.StandardClaims, error) {
	value, ok := c.Get("claims")
	if ok {
		claims, ok := value.(*jwt.StandardClaims)
		if ok {
			return claims, nil
		}
	}
	tokenString := c.GetHeader("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
		claims, err := utils.JwtDecodeToken(tokenString, config.Get().Key)
		if err == nil {
			c.Set("claims", claims)
		}
		return claims, err
	}
	return nil, errors.New("token不存在")
}

func AuthGetAccountID(c *gin.Context) string {
	claims, err := AuthDecodeToken(c)
	if err != nil {
		return ""
	}
	if claims == nil {
		return ""
	}
	return claims.Id
}

func AuthGetAccountOrError(c *gin.Context) (*bolejiang.Account, error) {
	value, exists := c.Get("account")
	if exists {
		account := value.(*bolejiang.Account)
		return account, nil
	}

	userId := AuthGetAccountID(c)
	var account bolejiang.Account
	ok, err := db.Default().Where("id = ?", userId).Get(&account)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("用户不存在")
	}
	if account.Status != 0 {
		return nil, errors.New("用户状态异常")
	}
	return &account, nil
}
