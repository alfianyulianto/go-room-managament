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

type RoomRepositoryImpl struct {
}

func NewRoomRepositoryImpl() *RoomRepositoryImpl {
	return &RoomRepositoryImpl{}
}

func (r RoomRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, filter web.RoomFilter) []domain.Room {
	query := "select rooms.id, rooms.room_category_id, room_categories.name, rooms.code, rooms.name, rooms.condition, rooms.note, rooms.created_at, rooms.updated_at from rooms join room_categories on room_categories.id = rooms.room_category_id"
	var args []interface{}
	var conditions []string

	if filter.Search != nil {
		conditions = append(conditions, "(rooms.name like ? or rooms.code like ? or rooms.note like ?)")
		args = append(args, "%"+*filter.Search+"%")
		args = append(args, "%"+*filter.Search+"%")
		args = append(args, "%"+*filter.Search+"%")
	}

	if filter.RoomCategoryId != nil {
		conditions = append(conditions, "rooms.room_category_id = ?")
		args = append(args, filter.RoomCategoryId)
	}

	if filter.Condition != nil {
		conditions = append(conditions, "rooms.condition = ?")
		args = append(args, *filter.Condition)
	}

	if len(conditions) > 0 {
		query += " where " + strings.Join(conditions, " and ")
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	halpers.IfPanicError(err)
	defer rows.Close()

	var rooms []domain.Room
	for rows.Next() {
		var room domain.Room
		err = rows.Scan(&room.Id, &room.RoomCategoryId, &room.RoomCategoryName, &room.Code, &room.Name, &room.Condition, &room.Note, &room.CreatedAt, &room.UpdatedAt)
		halpers.IfPanicError(err)
		rooms = append(rooms, room)
	}
	return rooms
}

func (r RoomRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (domain.Room, error) {
	rows, err := tx.QueryContext(ctx, "select rooms.id, rooms.room_category_id, room_categories.name, rooms.code, rooms.name, rooms.condition, rooms.note, rooms.created_at, rooms.updated_at from rooms join room_categories on room_categories.id = rooms.room_category_id where rooms.id = ?", id)
	halpers.IfPanicError(err)
	defer rows.Close()

	var room domain.Room
	if rows.Next() {
		err = rows.Scan(&room.Id, &room.RoomCategoryId, &room.RoomCategoryName, &room.Code, &room.Name, &room.Condition, &room.Note, &room.CreatedAt, &room.UpdatedAt)
		halpers.IfPanicError(err)
		return room, nil
	} else {
		return room, errors.New("room not found")
	}
}

func (r RoomRepositoryImpl) FindByName(ctx context.Context, dbOrTx QueryExecutor, name string) (*domain.Room, error) {
	rows, err := dbOrTx.QueryContext(ctx, "select * from rooms where name = ?", name)
	halpers.IfPanicError(err)
	defer rows.Close()

	var room domain.Room
	if rows.Next() {
		err = rows.Scan(&room.Id, &room.RoomCategoryId, &room.Code, &room.Name, &room.Condition, &room.Note, &room.CreatedAt, &room.UpdatedAt)
		halpers.IfPanicError(err)
		return &room, nil
	} else {
		return nil, errors.New("room not found")
	}
}

func (r RoomRepositoryImpl) FindByCode(ctx context.Context, dbOrTx QueryExecutor, code string) (*domain.Room, error) {
	rows, err := dbOrTx.QueryContext(ctx, "select * from rooms where code = ?", code)
	halpers.IfPanicError(err)
	defer rows.Close()

	var room domain.Room
	if rows.Next() {
		err = rows.Scan(&room.Id, &room.RoomCategoryId, &room.Code, &room.Name, &room.Condition, &room.Note, &room.CreatedAt, &room.UpdatedAt)
		halpers.IfPanicError(err)
		return &room, nil
	} else {
		return nil, errors.New("room not found")
	}
}

func (r RoomRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, room domain.Room) domain.Room {
	result, err := tx.ExecContext(ctx, "insert into rooms (room_category_id, name, code, `condition`, note ) values (?,?,?,?,?)", room.RoomCategoryId, room.Name, room.Code, room.Condition, room.Note)
	halpers.IfPanicError(err)

	//	get id
	id, err := result.LastInsertId()
	halpers.IfPanicError(err)

	rows, err := tx.QueryContext(ctx, "select rooms.id, rooms.room_category_id, room_categories.name as room_category_name, rooms.code, rooms.name, rooms.condition, rooms.note, rooms.created_at, rooms.updated_at from rooms join room_categories on room_categories.id = rooms.room_category_id where rooms.id = ?", id)
	halpers.IfPanicError(err)
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&room.Id, &room.RoomCategoryId, &room.RoomCategoryName, &room.Code, &room.Name, &room.Condition, &room.Note, &room.CreatedAt, &room.UpdatedAt)
		halpers.IfPanicError(err)
	}
	return room
}

func (r RoomRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, room domain.Room) domain.Room {
	_, err := tx.ExecContext(ctx, "update rooms set room_category_id = ?, name = ? , code = ?, `condition` = ?, note =? where id =?", room.RoomCategoryId, room.Name, room.Code, room.Condition, room.Note, room.Id)
	halpers.IfPanicError(err)

	//	get id
	rows, err := tx.QueryContext(ctx, "select rooms.id, rooms.room_category_id, room_categories.name, rooms.code, rooms.name, rooms.condition, rooms.note, rooms.created_at, rooms.updated_at from rooms join room_categories on room_categories.id = rooms.room_category_id where rooms.id = ?", room.Id)
	halpers.IfPanicError(err)
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&room.Id, &room.RoomCategoryId, &room.RoomCategoryName, &room.Code, &room.Name, &room.Condition, &room.Note, &room.CreatedAt, &room.UpdatedAt)
		halpers.IfPanicError(err)
	}
	return room
}

func (r RoomRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int64) {
	_, err := tx.ExecContext(ctx, "delete from rooms where id = ?", id)
	halpers.IfPanicError(err)
}
