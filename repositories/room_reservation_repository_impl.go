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

type RoomReservationRepositoryImpl struct {
}

func NewRoomReservationRepositoryImpl() *RoomReservationRepositoryImpl {
	return &RoomReservationRepositoryImpl{}
}

func (r RoomReservationRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, filter web.RoomReservationFilter) []domain.RoomReservation {
	query := "select room_reservations.id, room_reservations.room_id, rooms.name, room_reservations.user_id, users.name, room_reservations.start_at, room_reservations.end_at, room_reservations.purpose, room_reservations.status, room_reservations.approve_id, approver.name, room_reservations.file, room_reservations.created_at, room_reservations.updated_at from room_reservations join users on users.id = room_reservations.user_id join rooms on rooms.id = room_reservations.room_id left join users as approver on approver.id = room_reservations.approve_id"

	var args []interface{}
	var conditions []string

	if filter.Search != nil {
		conditions = append(conditions, "(users.name like ? or rooms.name like ? or room_reservations.purpose like ?)")
		args = append(args, "%"+*filter.Search+"%")
		args = append(args, "%"+*filter.Search+"%")
		args = append(args, "%"+*filter.Search+"%")
	}

	if filter.StartAt != nil {
		conditions = append(conditions, "room_reservations.start_at = ?")
		args = append(args, filter.StartAt)
	}

	if filter.EndAt != nil {
		conditions = append(conditions, "room_reservations.end_at = ?")
		args = append(args, *filter.EndAt)
	}

	if filter.Status != nil {
		conditions = append(conditions, "room_reservations.status = ?")
		args = append(args, *filter.Status)
	}

	if len(conditions) > 0 {
		query += " where " + strings.Join(conditions, " and ")
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	halpers.IfPanicError(err)
	defer rows.Close()

	var roomReservations []domain.RoomReservation
	for rows.Next() {
		var roomReservation domain.RoomReservation
		err = rows.Scan(&roomReservation.Id, &roomReservation.RoomId, &roomReservation.RoomName, &roomReservation.UserId, &roomReservation.UserName, &roomReservation.StartAt, &roomReservation.EndAt, &roomReservation.Purpose, &roomReservation.Status, &roomReservation.ApproveId, &roomReservation.ApproveName, &roomReservation.File, &roomReservation.CreatedAt, &roomReservation.UpdatedAt)
		halpers.IfPanicError(err)
		roomReservations = append(roomReservations, roomReservation)
	}
	return roomReservations
}

func (r RoomReservationRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (domain.RoomReservation, error) {
	rows, err := tx.QueryContext(ctx, "select room_reservations.id, room_reservations.user_id, room_reservations.room_id, room_reservations.start_at, room_reservations.end_at, room_reservations.purpose, room_reservations.status, room_reservations.approve_id, room_reservations.file, room_reservations.created_at, room_reservations.updated_at from room_reservations where id=?", id)
	halpers.IfPanicError(err)
	defer rows.Close()

	var roomReservation domain.RoomReservation
	if rows.Next() {
		err = rows.Scan(&roomReservation.Id, &roomReservation.UserId, &roomReservation.RoomId, &roomReservation.StartAt, &roomReservation.EndAt, &roomReservation.Purpose, &roomReservation.Status, &roomReservation.ApproveId, &roomReservation.File, &roomReservation.CreatedAt, &roomReservation.UpdatedAt)
		halpers.IfPanicError(err)
		return roomReservation, nil
	} else {
		return roomReservation, errors.New("room reservation not found")
	}
}

func (r RoomReservationRepositoryImpl) HasOverlap(ctx context.Context, dbOrTx QueryExecutor, roomId int64, startAt string, endAt string) *domain.RoomReservation {
	rows, err := dbOrTx.QueryContext(ctx, "select id from room_reservations where room_id=? and start_at <= ? and end_at >=?", roomId, startAt, endAt)
	halpers.IfPanicError(err)
	defer rows.Close()

	var reservation domain.RoomReservation
	if rows.Next() {
		err = rows.Scan(&reservation.Id)
		halpers.IfPanicError(err)
		return &reservation
	} else {
		return nil
	}
}

func (r RoomReservationRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, roomReservation domain.RoomReservation) domain.RoomReservation {
	result, err := tx.ExecContext(ctx, "insert into room_reservations (user_id, room_id, start_at, end_at, purpose, status, approve_id, file) values (?,?,?,?,?,?,?,?)", roomReservation.UserId, roomReservation.RoomId, roomReservation.StartAt, roomReservation.EndAt, roomReservation.Purpose, roomReservation.Status, roomReservation.ApproveId, roomReservation.File)
	halpers.IfPanicError(err)

	//	get id
	id, err := result.LastInsertId()
	halpers.IfPanicError(err)

	roomReservation, err = r.FindById(ctx, tx, id)
	halpers.IfPanicError(err)
	return roomReservation
}

func (r RoomReservationRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, roomReservation domain.RoomReservation) domain.RoomReservation {
	_, err := tx.ExecContext(ctx, "update room_reservations set user_id=?, room_id=?, start_at=?, end_at=?, purpose=?, approve_id=?, file=? where id=?", roomReservation.UserId, roomReservation.RoomId, roomReservation.StartAt, roomReservation.EndAt, roomReservation.Purpose, roomReservation.ApproveId, roomReservation.File, roomReservation.Id)
	halpers.IfPanicError(err)

	roomReservation, err = r.FindById(ctx, tx, roomReservation.Id)
	halpers.IfPanicError(err)
	return roomReservation
}

func (r RoomReservationRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int64) {
	_, err := tx.ExecContext(ctx, "delete from room_reservations where id=?", id)
	halpers.IfPanicError(err)
}
