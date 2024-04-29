package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"strings"
	"time"
)

const TokenExpireDuration = time.Hour * 24

type MyClaims struct {
	AccountBo bo.AccountBo `json:"account"`
	jwt.StandardClaims
}

func GenToken(accountBo bo.AccountBo) (string, error) {
	c := MyClaims{
		AccountBo: accountBo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "h-ui",
		},
	}
	config, err := dao.GetConfig("key = ?", constant.JwtSecret)
	if err != nil {
		return "", errors.New(constant.SysError)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(*config.Value))
}

func ParseToken(tokenString string) (*MyClaims, error) {
	config, err := dao.GetConfig("key = ?", constant.JwtSecret)
	if err != nil {
		return nil, errors.New(constant.SysError)
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(*config.Value), nil
	})
	if err != nil {
		return nil, errors.New(constant.IllegalTokenError)
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New(constant.TokenExpiredError)
}

func GetToken(c *gin.Context) string {
	tokenStr := c.Request.Header.Get("Authorization")
	if tokenStr == "" {
		return ""
	}
	return strings.SplitN(tokenStr, " ", 2)[1]
}
