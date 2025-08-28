package repositories

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/web"
)

type RoomCategoryRepositoryImpl struct {
}

func NewRoomCategoryRepositoryImp(db *sql.DB) *RoomCategoryRepositoryImpl {
	return &RoomCategoryRepositoryImpl{}
}

func (r *RoomCategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, filter web.RoomCategoryFilter) []domain.RoomCategory {
	query := "select id, name, created_at, updated_at from room_categories"
	var args []interface{}
	var conditions []string

	if filter.Search != nil {
		conditions = append(conditions, "name like ?")
		args = append(args, "%"+*filter.Search+"%")
	}

	if len(conditions) > 0 {
		query += " where " + strings.Join(conditions, " and ")
	}
	rows, err := tx.QueryContext(ctx, query, args...)
	halpers.IfPanicError(err)
	defer rows.Close()

	var roomCategories []domain.RoomCategory

	for rows.Next() {
		var roomCategory domain.RoomCategory
		err = rows.Scan(&roomCategory.Id, &roomCategory.Name, &roomCategory.CreatedAt, &roomCategory.UpdatedAt)
		halpers.IfPanicError(err)

		roomCategories = append(roomCategories, roomCategory)
	}
	return roomCategories
}

func (r *RoomCategoryRepositoryImpl) FindById(ctx context.Context, dbOrTx QueryExecutor, id int64) (*domain.RoomCategory, error) {
	rows, err := dbOrTx.QueryContext(ctx, "select * from room_categories where id = ?", id)
	halpers.IfPanicError(err)
	defer rows.Close()

	var roomCategory domain.RoomCategory
	if rows.Next() {
		err = rows.Scan(&roomCategory.Id, &roomCategory.Name, &roomCategory.CreatedAt, &roomCategory.UpdatedAt)
		halpers.IfPanicError(err)
		return &roomCategory, nil
	} else {
		return nil, errors.New("room category not found")
	}
}

func (r *RoomCategoryRepositoryImpl) FindByName(ctx context.Context, dbOrTx QueryExecutor, name string) (*domain.RoomCategory, error) {
	rows, err := dbOrTx.QueryContext(ctx, "select * from room_categories where name = ? limit 1", name)
	halpers.IfPanicError(err)
	defer rows.Close()

	var roomCategory domain.RoomCategory
	if rows.Next() {
		err = rows.Scan(&roomCategory.Id, &roomCategory.Name, &roomCategory.CreatedAt, &roomCategory.UpdatedAt)
		halpers.IfPanicError(err)
		return &roomCategory, nil
	} else {
		return nil, errors.New("room category not found")
	}
}

func (r *RoomCategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, roomCategory domain.RoomCategory) domain.RoomCategory {
	result, err := tx.ExecContext(ctx, "insert into room_categories (name) values (?)", roomCategory.Name)
	halpers.IfPanicError(err)

	//	get id
	id, err := result.LastInsertId()
	halpers.IfPanicError(err)

	//find by id
	rows, err := tx.QueryContext(ctx, "select id, name, created_at, updated_at from room_categories where id = ?", id)
	halpers.IfPanicError(err)
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&roomCategory.Id, &roomCategory.Name, &roomCategory.CreatedAt, &roomCategory.UpdatedAt)
		halpers.IfPanicError(err)
	}

	return roomCategory
}

func (r *RoomCategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, roomCategory domain.RoomCategory) domain.RoomCategory {
	_, err := tx.ExecContext(ctx, "update room_categories set name = ? where id = ?", roomCategory.Name, roomCategory.Id)
	halpers.IfPanicError(err)

	//find by id
	rows, err := tx.QueryContext(ctx, "select * from room_categories where id = ?", roomCategory.Id)
	halpers.IfPanicError(err)
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&roomCategory.Id, &roomCategory.Name, &roomCategory.CreatedAt, &roomCategory.UpdatedAt)
		halpers.IfPanicError(err)
	}

	return roomCategory
}

func (r *RoomCategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int64) {
	_, err := tx.ExecContext(ctx, "delete from room_categories where id = ?", id)
	halpers.IfPanicError(err)
}
