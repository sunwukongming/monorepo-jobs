/**

Filename: 		jwt.go
Author: 		gengming - gengming.zb@ccbft.com
Description:	loongrpc jwt logic
Create:			2022-07-01 15:03:37
Last Modified:	2022-07-01 16:32:42

*/

package utils

import (
	"app/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GetToken(accountId int) (string, error) {
	ts := time.Now().Unix() / 1800 * 1800
	tokenString, err := JwtBuildToken(jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: ts + 7*24*60*60*100,
		Id:        strconv.Itoa(accountId),
		IssuedAt:  ts,
		Issuer:    "",
		NotBefore: ts,
		Subject:   "",
	}, config.Get().Key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func JwtBuildToken(claims jwt.StandardClaims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func JwtDecodeToken(tokenString string, key string) (*jwt.StandardClaims, error) {
	var claims jwt.StandardClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, nil
}
