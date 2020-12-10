package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const SecretKey = "MySuperSafeSecretKey"

type CustomJWTClaim struct {
	UserId int64
	jwt.StandardClaims
}

func GenerateJWT() (tokenString string, expireTime time.Time) {
	maxAge := 60 * 60 * 24
	expireTime = time.Now().Add(time.Duration(maxAge) * time.Second)
	myCustomClaim := &CustomJWTClaim{
		UserId: 6,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "RalXYZ",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myCustomClaim)
	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		panic(err)
	}
	// fmt.Printf("The generated Token is: %s\n", tokenString)
	return
}

func ParseToken(tokenString string) (*CustomJWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomJWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpexted singing method: %v\n", token.Header["alg"])
		} else {
			return []byte(SecretKey), nil
		}
	})
	if claims, ok := token.Claims.(*CustomJWTClaim); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func SetCookie(c *echo.Context, name string, token string, expireTime *time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = *expireTime
	(*c).SetCookie(cookie)
}

func GetCookie(c *echo.Context, name string) (token string, err error) {
	cookie, err := (*c).Cookie(name)
	if err != nil {
		return
	}
	token = cookie.Value
	return
}

func Authentication(c *echo.Context, name string) bool {
	if token, err := GetCookie(c, name); err != nil {
		return false
	} else if _, err = ParseToken(token); err != nil {
		return false
	} else {
		return true
	}
}
