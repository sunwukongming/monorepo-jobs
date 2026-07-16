package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var key = "YOU7ASY0OevEJrQg"

func TestJwtBuildToken(t *testing.T) {
	tokenString, err := JwtBuildToken(jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: time.Now().Unix(),
		Id:        "1",
		IssuedAt:  time.Now().Unix(),
		Issuer:    "",
		NotBefore: time.Now().Unix(),
		Subject:   "",
	}, key)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	t.Log(tokenString)
}

func TestJwtDecodeToken(t *testing.T) {
	tokenString, err := JwtBuildToken(jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: time.Now().Unix() + 86400,
		Id:        "1",
		IssuedAt:  time.Now().Unix(),
		Issuer:    "",
		NotBefore: time.Now().Unix(),
		Subject:   "",
	}, key)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	claims, err := JwtDecodeToken(tokenString, key)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else {
		if claims.Id != "1" {
			t.Fail()
		}
		t.Log(claims)
	}

}
