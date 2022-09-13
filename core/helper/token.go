package helper

import (
	"cloud-disk/core/define"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 生成token
func GenerateToken(id int,identity,name string,expireTime int) (string,error){
	uc := define.UserClaim{
		Id: id,
		Identity: identity,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expireTime)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,uc)
	jwtToken,err := token.SignedString([]byte(define.JwtKey))
	if err!=nil{
		return "", err
	}
	return jwtToken,nil
}

// 解析token
func ParseToken(token string)(*define.UserClaim,error){
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})

	if err!=nil{
		return nil,err
	}

	if !claims.Valid{
		return uc,errors.New("token is invalid")
	}

	return uc,err
}