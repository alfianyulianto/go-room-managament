package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/domain"
)

type RoomImageRepositoryImpl struct {
}

func NewRoomImageRepositoryImpl() *RoomImageRepositoryImpl {
	return &RoomImageRepositoryImpl{}
}

func (r RoomImageRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, roomId int64) []domain.RoomImage {
	rows, err := tx.QueryContext(ctx, "select * from room_images where room_id=?", roomId)
	halpers.IfPanicError(err)

	var roomImages []domain.RoomImage
	for rows.Next() {
		var roomImage domain.RoomImage
		err = rows.Scan(&roomImage.Id, &roomImage.RoomId, &roomImage.Path, &roomImage.CreatedAt, &roomImage.UpdatedAt)
		halpers.IfPanicError(err)

		roomImages = append(roomImages, roomImage)
	}
	return roomImages
}

func (r RoomImageRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, roomId int64, id int64) (domain.RoomImage, error) {
	rows, err := tx.QueryContext(ctx, "select * from room_images where room_id = ? and id =?", roomId, id)
	halpers.IfPanicError(err)
	defer rows.Close()

	var roomImage domain.RoomImage
	if rows.Next() {
		err = rows.Scan(&roomImage.Id, &roomImage.RoomId, &roomImage.Path, &roomImage.CreatedAt, &roomImage.UpdatedAt)
		halpers.IfPanicError(err)
		return roomImage, nil
	} else {
		return roomImage, errors.New("image found")
	}
}

func (r RoomImageRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, roomImage domain.RoomImage) domain.RoomImage {
	result, err := tx.ExecContext(ctx, "insert into room_images (room_id, path) values (?,?)", roomImage.RoomId, roomImage.Path)
	halpers.IfPanicError(err)

	id, err := result.LastInsertId()

	//	get by id
	rows, err := tx.QueryContext(ctx, "select * from room_images where id = ?", id)
	halpers.IfPanicError(err)
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&roomImage.Id, &roomImage.RoomId, &roomImage.Path, &roomImage.CreatedAt, &roomImage.UpdatedAt)
		halpers.IfPanicError(err)
	}

	return roomImage
}

func (r RoomImageRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, roomImage domain.RoomImage) domain.RoomImage {
	_, err := tx.ExecContext(ctx, "update room_images set path => ? where id = ?", roomImage.Path, roomImage.Id)
	halpers.IfPanicError(err)

	//get id
	rows, err := tx.QueryContext(ctx, "select * from room_images where id = ?", roomImage.Id)
	halpers.IfPanicError(err)
	if rows.Next() {
		err = rows.Scan(&roomImage.Id, &roomImage.RoomId, &roomImage.Path, &roomImage.CreatedAt, &roomImage.UpdatedAt)
		halpers.IfPanicError(err)
	}
	return roomImage
}

func (r RoomImageRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, roomId int64, id int64) {
	_, err := tx.ExecContext(ctx, "delete from room_images where id = ?", id)
	halpers.IfPanicError(err)
}
