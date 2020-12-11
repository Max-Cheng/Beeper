package common

import (
	"Beeper/backend/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	Userid uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string,error) {
	expirasionTime := time.Now().Add(24*time.Hour)
	claims :=&Claims{
		Userid: user.ID,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expirasionTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer:	"Cola.Inc",
			Subject:"user token",
		},
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err:=token.SignedString(jwtKey)
	if err != nil {
		return "",err
	}
	return tokenString,nil
}
func ParseToken(tokenstring string)	(*jwt.Token,*Claims,error) {
	claims :=&Claims{}

	token,err:=jwt.ParseWithClaims(tokenstring,claims,func(token *jwt.Token) (interface{}, error) {
		return jwtKey,nil
	})

	return token,claims,err
}