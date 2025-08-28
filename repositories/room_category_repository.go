package repositories

import (
	"context"
	"database/sql"

	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type QueryExecutor interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type RoomCategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, filter web.RoomCategoryFilter) []domain.RoomCategory
	FindById(ctx context.Context, dbOrTx QueryExecutor, id int64) (*domain.RoomCategory, error)
	FindByName(ctx context.Context, dbOrTx QueryExecutor, name string) (*domain.RoomCategory, error)
	Save(ctx context.Context, tx *sql.Tx, roomCategory domain.RoomCategory) domain.RoomCategory
	Update(ctx context.Context, tx *sql.Tx, roomCategory domain.RoomCategory) domain.RoomCategory
	Delete(ctx context.Context, tx *sql.Tx, id int64)
}
