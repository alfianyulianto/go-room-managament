package services

import (
	"context"

	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type RoomService interface {
	FindAll(ctx context.Context, filter web.RoomFilter) []web.RoomResponse
	FindById(ctx context.Context, id int64) web.RoomResponse
	Create(ctx context.Context, request request.RoomCreateRequest) web.RoomResponse
	Update(ctx context.Context, request request.RoomUpdateRequest) web.RoomResponse
	Delete(ctx context.Context, id int64)
}
