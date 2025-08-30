package repositories

import (
	"context"
	"database/sql"

	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type RoomReservationRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, filter web.RoomReservationFilter) []domain.RoomReservation
	FindById(ctx context.Context, tx *sql.Tx, id int64) (domain.RoomReservation, error)
	HasOverlap(ctx context.Context, dbOrTx QueryExecutor, roomId int64, startAt string, endAt string) *domain.RoomReservation
	Save(ctx context.Context, tx *sql.Tx, roomReservation domain.RoomReservation) domain.RoomReservation
	Update(ctx context.Context, tx *sql.Tx, roomReservation domain.RoomReservation) domain.RoomReservation
	Delete(ctx context.Context, tx *sql.Tx, id int64)
	Accepted(ctx context.Context, tx *sql.Tx, roomReservation domain.RoomReservation) domain.RoomReservation
	Rejected(ctx context.Context, tx *sql.Tx, roomReservation domain.RoomReservation) domain.RoomReservation
}
