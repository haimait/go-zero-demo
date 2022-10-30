package jwtToken

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaim struct {
	Id   int
	Name string
	jwt.StandardClaims
}

var JwtKey = "user-demo-key"

func GenerateToken(id int, name string, second int) (string, error) {
	uc := UserClaim{
		Id:   id,
		Name: name,
		//设置有效期 多少秒内有效
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyzeToken
// Token 解析
func AnalyzeToken(token string) (*UserClaim, error) {
	uc := new(UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}
