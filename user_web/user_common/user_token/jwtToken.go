package user_token

import (
	"github.com/golang-jwt/jwt"
	"py_gomall/v2/user_web/user_config"
	"time"
)

type Claims struct {
	Mobile string `json:"mobile"`
	jwt.StandardClaims
}

func Release(mobile string) (string, error) {
	jwtKey := []byte(user_config.AppConf.Jwt.Key)
	issuer := user_config.AppConf.Jwt.Issuer
	subject := user_config.AppConf.Jwt.Subject
	expireTimeMinute := time.Duration(user_config.AppConf.Jwt.ExpireMinute) * time.Minute
	expireTime := time.Now().Add(expireTimeMinute)
	claims := &Claims{
		mobile,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
			Subject:   subject,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Parse(tokenStr string) (*jwt.Token, *Claims, error) {
	jwtKey := []byte(user_config.AppConf.Jwt.Key)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
