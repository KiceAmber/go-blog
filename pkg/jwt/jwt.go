package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenExpireDuration token的过期时间
const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("Amber-golang")

type MyClaims struct {
	UserID   int64
	Username string
	jwt.StandardClaims
}

func GenToken(userID int64, username string) (string, error) {
	// 首先创建一个自己声明的 claims 对象
	c := MyClaims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "amber-blog",
		},
	}

	// 指定签名方法来创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名，并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析 token
func ParseToken(tokenString string) (*MyClaims, error) {
	var myClaims = new(MyClaims)

	token, err := jwt.ParseWithClaims(tokenString, myClaims, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return myClaims, nil
	}
	return nil, errors.New("invalid token")
}
