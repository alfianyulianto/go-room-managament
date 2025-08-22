package services

import (
	"context"

	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type UserService interface {
	FindAll(ctx context.Context, filter web.UserFilter) []web.UserResponse
	FindById(ctx context.Context, id int64) web.UserResponse
	Create(ctx context.Context, request request.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request request.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, id int64)
}
