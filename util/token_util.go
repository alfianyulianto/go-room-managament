package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/alfianyulianto/go-room-managament/config"
	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/domain"
)

type TokenUtil struct {
	SecretKey string
}

func NewTokenUtil() *TokenUtil {
	return &TokenUtil{SecretKey: config.Cfg.SecretKey}
}

func (t TokenUtil) CreateToken(auth domain.Auth) (string, error) {
	expire := time.Now().Add(time.Hour * 24).UnixMilli()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     auth.Id,
		"name":   auth.Name,
		"email":  auth.Email,
		"phone":  auth.Phone,
		"level":  auth.Level,
		"expire": expire,
	})

	jwtToken, err := token.SignedString([]byte(t.SecretKey))
	halpers.IfPanicError(err)
	return jwtToken, nil
}

func (t TokenUtil) ParseToken(jwtToken string) (*domain.Auth, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})
	if err != nil {
		return nil, errors.New("token is invalid")
	}

	claims := token.Claims.(jwt.MapClaims)

	expire := claims["expire"].(float64)

	if int64(expire) < time.Now().UnixMilli() {
		return nil, errors.New("token is expired")
	}

	id := claims["id"].(float64)
	name := claims["name"].(string)
	email := claims["email"].(string)
	phone := claims["phone"].(string)
	level := claims["level"].(string)

	return &domain.Auth{
		Id:    int64(id),
		Name:  name,
		Email: email,
		Phone: phone,
		Level: level,
	}, nil
}
