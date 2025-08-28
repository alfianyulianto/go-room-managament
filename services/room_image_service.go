package services

import (
	"context"

	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type RoomImageService interface {
	FindAll(ctx context.Context, roomId int64) []web.RoomImageResponse
	FindById(ctx context.Context, roomId int64, id int64) web.RoomImageResponse
	Create(ctx context.Context, request request.RoomImageCreateRequest) web.RoomImageResponse
	Update(ctx context.Context, request request.RoomImageUpdateRequest) web.RoomImageResponse
	Delete(ctx context.Context, roomId int64, id int64)
}
