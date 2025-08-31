package services

import (
	"context"

	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type RoomReservationService interface {
	FindAll(ctx context.Context, filter web.RoomReservationFilter) []web.RoomReservationResponse
	FindById(ctx context.Context, id int64) web.RoomReservationResponse
	Create(ctx context.Context, request request.RoomReservationCreateRequest) web.RoomReservationResponse
	Update(ctx context.Context, request request.RoomReservationUpdateRequest) web.RoomReservationResponse
	Delete(ctx context.Context, id int64)
	Accepted(ctx context.Context, id int64) web.RoomReservationResponse
	Rejected(ctx context.Context, id int64) web.RoomReservationResponse
}
