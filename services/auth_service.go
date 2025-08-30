package services

import (
	"context"

	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type AuthService interface {
	Login(ctx context.Context, request request.LoginRequest) (*web.TokenResponse, error)
	Register(ctx context.Context, request request.RegisterRequest) web.UserResponse
}
