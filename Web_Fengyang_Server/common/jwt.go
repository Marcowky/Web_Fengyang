package common

import (
	"Web_Fengyang_Server/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt加密密钥
var jwtKey = []byte("a_secret_key")

type Claims struct {
	UserId uint
	UserType string
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) { //创建token
	// token的有效期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		// 自定义字段
		UserId: user.ID,
		UserType: user.UserType,
		// 标准字段
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expirationTime.Unix(),
			// 发放时间
			IssuedAt: time.Now().Unix(),
		},
	}
	// 使用jwt包提供的NewWithClaims函数创建一个新的token变量，其中第一个参数是加密方法（HS256表示采用HMAC-SHA256加密算法），第二个参数是claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	// 返回token
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) { //解析token
	claims := &Claims{}
	//使用jwt包提供的ParseWithClaims函数解析token，该函数会将token解密并验证签名。如果解密和验证成功，将返回一个jwt.Token类型的token和Claims类型的claims。第三个参数是一个回调函数，用于在解密和验证过程中使用密钥。回调函数返回的是密钥。
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
