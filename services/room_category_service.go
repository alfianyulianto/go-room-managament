package services

import (
	"context"

	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type RoomCategoryService interface {
	FindAll(ctx context.Context, filter web.RoomCategoryFilter) []web.RoomCategoryResponse
	FindById(ctx context.Context, id int64) web.RoomCategoryResponse
	Create(ctx context.Context, request request.RoomCategoryCreateRequest) web.RoomCategoryResponse
	Update(ctx context.Context, request request.RoomCategoryUpdateRequest) web.RoomCategoryResponse
	Delete(ctx context.Context, id int64)
}
