package repositories

import (
	"context"
	"database/sql"

	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type RoomRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, filter web.RoomFilter) []domain.Room
	FindById(ctx context.Context, dbOrTx QueryExecutor, id int64) (*domain.Room, error)
	FindByName(ctx context.Context, dbOrTx QueryExecutor, name string) (*domain.Room, error)
	FindByCode(ctx context.Context, dbOrTx QueryExecutor, code string) (*domain.Room, error)
	Save(ctx context.Context, tx *sql.Tx, room domain.Room) domain.Room
	Update(ctx context.Context, tx *sql.Tx, room domain.Room) domain.Room
	Delete(ctx context.Context, tx *sql.Tx, id int64)
}
