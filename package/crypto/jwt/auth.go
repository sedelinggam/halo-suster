package cryptoJWT

import (
	"halo-suster/config"
	"halo-suster/package/lumen"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type key int

const (
	KeyPhoneNumber key = iota
)

type JWTClaims struct {
	Id          string
	PhoneNumber string
	jwt.RegisteredClaims
}

func GenerateToken(id, phoneNumber string) (*string, error) {
	secret := config.JWTSecret()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Id:          id,
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
		},
	})

	tokenString, err := token.SignedString([]byte(secret))
	return &tokenString, err
}

type JWTPayload struct {
	Id          string
	PhoneNumber string
}

func VerifyToken(token string) (*JWTPayload, error) {
	secret := config.JWTSecret()

	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims.RegisteredClaims.ExpiresAt.Before(time.Now()) {
		return nil, err
	}

	payload := &JWTPayload{
		Id:          claims.Id,
		PhoneNumber: claims.PhoneNumber,
	}

	return payload, nil
}

func JWTConfig() echojwt.Config {
	config := echojwt.Config{
		SigningKey: []byte(config.JWTSecret()),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTClaims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return lumen.FromError(lumen.NewError(lumen.ErrUnauthorized, err)).SendResponse(c)
		},
	}
	return config
}
