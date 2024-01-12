package encrypt

import (
	"WePanel/orm"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("rtdbskofd")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func GetToken(user orm.User) (string, error) {
	//设置过期时间：2小时
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{ //创建Claims实例
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //根据claims创建token
	if tokenString, err := token.SignedString(jwtKey); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (i interface{}, err error) {
			return jwtKey, nil
		})
	return token, claims, err
}
