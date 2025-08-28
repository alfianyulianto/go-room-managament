package repositories

import (
	"context"
	"database/sql"

	"github.com/alfianyulianto/go-room-managament/model/domain"
)

type RoomImageRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, roomId int64) []domain.RoomImage
	FindById(ctx context.Context, tx *sql.Tx, roomId int64, id int64) (domain.RoomImage, error)
	Save(ctx context.Context, tx *sql.Tx, roomImage domain.RoomImage) domain.RoomImage
	Update(ctx context.Context, tx *sql.Tx, roomImage domain.RoomImage) domain.RoomImage
	Delete(ctx context.Context, tx *sql.Tx, roomId int64, id int64)
}
