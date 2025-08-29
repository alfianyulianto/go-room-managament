package repositories

import (
	"context"
	"database/sql"

	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, filter web.UserFilter) []domain.User
	FindById(ctx context.Context, dbOrTx QueryExecutor, id int64) (*domain.User, error)
	FindByName(ctx context.Context, dbOrTx QueryExecutor, name string) (*domain.User, error)
	FindByEmail(ctx context.Context, dbOrTx QueryExecutor, email string) (*domain.User, error)
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, id int64)
}
