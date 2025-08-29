package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/alfianyulianto/go-room-managament/config"
	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
	"github.com/alfianyulianto/go-room-managament/repositories"
	"github.com/alfianyulianto/go-room-managament/storage"
)

type RoomReservationServiceImpl struct {
	repositories.RoomReservationRepository
	repositories.UserRepository
	repositories.RoomRepository
	*sql.DB
	*validator.Validate
	storage.FileStorage
}

func NewRoomReservationServiceImpl(roomReservationRepository repositories.RoomReservationRepository, userRepository repositories.UserRepository, roomRepository repositories.RoomRepository, DB *sql.DB, validate *validator.Validate, fileStorage storage.FileStorage) *RoomReservationServiceImpl {
	err := validate.RegisterValidation("existUser", func(fl validator.FieldLevel) bool {
		userId := fl.Field().Int()
		user, _ := userRepository.FindById(context.Background(), DB, userId)
		if user != nil {
			return true
		}
		return false
	})

	err = validate.RegisterValidation("existRoom", func(fl validator.FieldLevel) bool {
		roomId := fl.Field().Int()
		room, _ := roomRepository.FindById(context.Background(), DB, roomId)
		if room != nil {
			return true
		}
		return false
	})

	err = validate.RegisterValidation("file", func(fl validator.FieldLevel) bool {
		file, ok := fl.Field().Interface().(multipart.FileHeader)
		if !ok {
			return false
		}

		if file.Size > 5<<20 { // max 5MB
			return false
		}

		ext := strings.ToLower(filepath.Ext(file.Filename))
		allowedExts := strings.Split(fl.Param(), ";")

		for _, allowed := range allowedExts {
			if ext == "."+strings.ToLower(strings.TrimSpace(allowed)) {
				return true
			}
		}

		return false
	})

	err = validate.RegisterValidation("after", func(fl validator.FieldLevel) bool {
		fieldName := fl.Param()

		currentStr, ok1 := fl.Field().Interface().(string)
		if !ok1 || currentStr == "" {
			return false
		}

		otherField := fl.Parent().FieldByName(fieldName)
		if !otherField.IsValid() {
			return false
		}

		otherStr, ok2 := otherField.Interface().(string)
		if !ok2 || otherStr == "" {
			return false
		}

		current, err1 := time.Parse("2006-01-02 15:04:05", currentStr)
		other, err2 := time.Parse("2006-01-02 15:04:05", otherStr)
		if err1 != nil || err2 != nil {
			return false
		}

		return current.After(other)
	})

	err = validate.RegisterValidation("afterToday", func(fl validator.FieldLevel) bool {
		startStr, ok := fl.Field().Interface().(string)
		if !ok || startStr == "" {
			return false
		}

		start, err := time.Parse("2006-01-02 15:04:05", startStr)
		if err != nil {
			return false
		}

		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		return start.After(today)
	})
	err = validate.RegisterValidation("overlap", func(fl validator.FieldLevel) bool {
		startStr, ok := fl.Field().Interface().(string)
		if !ok || startStr == "" {
			return false
		}

		endStr := fl.Parent().FieldByName("EndAt").Interface().(string)
		if endStr == "" {
			return false
		}

		id := fl.Parent().FieldByName("Id").Interface().(int64)

		roomId := fl.Parent().FieldByName("RoomId").Interface().(int64)

		overlap := roomReservationRepository.HasOverlap(context.Background(), DB, roomId, startStr, endStr)
		fmt.Println(overlap.Id)
		fmt.Println(id)

		if overlap != nil {
			if overlap.Id == id {
				return true
			}
			return false
		}

		return true
	})
	halpers.IfPanicError(err)

	return &RoomReservationServiceImpl{RoomReservationRepository: roomReservationRepository, UserRepository: userRepository, RoomRepository: roomRepository, DB: DB, Validate: validate, FileStorage: fileStorage}
}

func (r RoomReservationServiceImpl) FindAll(ctx context.Context, filter web.RoomReservationFilter) []web.RoomReservationResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	reservations := r.RoomReservationRepository.FindAll(ctx, tx, filter)
	var roomReservationResponses []web.RoomReservationResponse
	for _, reservation := range reservations {
		halpers.IfPanicError(err)

		roomReservationResponses = append(roomReservationResponses, web.RoomReservationResponse{
			Id:        reservation.Id,
			UserId:    reservation.UserId,
			UserName:  &reservation.UserName,
			RoomId:    reservation.RoomId,
			RoomName:  reservation.RoomName,
			StarAt:    reservation.StartAt.Format("2006-01-02 15:04:05"),
			EndAt:     reservation.EndAt.Format("2006-01-02 15:04:05"),
			Purpose:   reservation.Purpose,
			Status:    reservation.Status,
			ApproveId: halpers.ConvertToInt64(reservation.ApproveId),
			File:      path.Join(config.Cfg.BaseUrl, reservation.File),
			CreatedAt: reservation.CreatedAt,
			UpdatedAt: reservation.UpdatedAt,
		})
	}
	return roomReservationResponses
}

func (r RoomReservationServiceImpl) FindById(ctx context.Context, id int64) web.RoomReservationResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	reservation, err := r.RoomReservationRepository.FindById(ctx, tx, id)
	halpers.IfPanicError(err)

	user, err := r.UserRepository.FindById(ctx, tx, reservation.UserId)
	halpers.IfPanicError(err)

	room, err := r.RoomRepository.FindById(ctx, tx, reservation.RoomId)
	halpers.IfPanicError(err)

	var approveResponse *web.UserResponse
	if reservation.ApproveId.Valid {
		approve, err := r.UserRepository.FindById(ctx, tx, reservation.ApproveId.Int64)
		fmt.Println(reservation.ApproveId)
		halpers.IfPanicError(err)

		if approve != nil {
			approveResponse = &web.UserResponse{
				Id:        approve.Id,
				Name:      approve.Name,
				Email:     approve.Email,
				Phone:     approve.Phone,
				CreatedAt: approve.CreatedAt,
				UpdatedAt: approve.UpdatedAt,
			}
		}
	}

	return web.RoomReservationResponse{
		Id:        reservation.Id,
		UserId:    reservation.UserId,
		RoomId:    reservation.RoomId,
		StarAt:    reservation.StartAt.Format("2006-01-02 15:04:05"),
		EndAt:     reservation.EndAt.Format("2006-01-02 15:04:05"),
		Purpose:   reservation.Purpose,
		Status:    reservation.Status,
		ApproveId: halpers.ConvertToInt64(reservation.ApproveId),
		File:      path.Join(config.Cfg.BaseUrl, reservation.File),
		CreatedAt: reservation.CreatedAt,
		UpdatedAt: reservation.UpdatedAt,
		User: &web.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Room: &web.RoomResponse{
			Id:             room.Id,
			RoomCategoryId: room.RoomCategoryId,
			Code:           room.Code,
			Name:           room.Name,
			Condition:      room.Condition,
			CreatedAt:      room.CreatedAt,
			UpdatedAt:      room.UpdatedAt,
		},
		Approve: approveResponse,
	}

}

func (r RoomReservationServiceImpl) Create(ctx context.Context, request request.RoomReservationCreateRequest) web.RoomReservationResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	err = r.Validate.Struct(request)
	halpers.IfPanicError(err)

	filePath, err := r.FileStorage.SaveFile(request.File, "room_reservations")
	halpers.IfPanicError(err)

	startAt, err := time.Parse("2006-01-02 15:04:05", request.StartAt)
	halpers.IfPanicError(err)

	endAt, err := time.Parse("2006-01-02 15:04:05", request.EndAt)
	halpers.IfPanicError(err)

	reservation := r.RoomReservationRepository.Save(ctx, tx, domain.RoomReservation{
		UserId:    request.UserId,
		RoomId:    request.RoomId,
		StartAt:   startAt,
		EndAt:     endAt,
		Purpose:   request.Purpose,
		Status:    "Pengajuan",
		ApproveId: sql.NullInt64{Valid: false},
		File:      filePath,
	})

	return web.RoomReservationResponse{
		Id:        reservation.Id,
		UserId:    reservation.UserId,
		RoomId:    reservation.RoomId,
		StarAt:    reservation.StartAt.Format("2006-01-02 15:04:05"),
		EndAt:     reservation.EndAt.Format("2006-01-02 15:04:05"),
		Purpose:   reservation.Purpose,
		Status:    reservation.Status,
		ApproveId: halpers.ConvertToInt64(reservation.ApproveId),
		File:      path.Join(config.Cfg.BaseUrl, reservation.File),
		CreatedAt: reservation.CreatedAt,
		UpdatedAt: reservation.UpdatedAt,
	}
}

func (r RoomReservationServiceImpl) Update(ctx context.Context, request request.RoomReservationUpdateRequest) web.RoomReservationResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	err = r.Validate.Struct(request)
	halpers.IfPanicError(err)

	filePath, err := r.FileStorage.SaveFile(request.File, "room_reservations")
	halpers.IfPanicError(err)

	startAt, err := time.Parse("2006-01-02 15:04:05", request.StartAt)
	halpers.IfPanicError(err)

	endAt, err := time.Parse("2006-01-02 15:04:05", request.EndAt)
	halpers.IfPanicError(err)

	reservation := r.RoomReservationRepository.Update(ctx, tx, domain.RoomReservation{
		Id:      request.Id,
		UserId:  request.UserId,
		RoomId:  request.RoomId,
		StartAt: startAt,
		EndAt:   endAt,
		Purpose: request.Purpose,
		File:    filePath,
	})

	return web.RoomReservationResponse{
		Id:        reservation.Id,
		UserId:    reservation.UserId,
		RoomId:    reservation.RoomId,
		StarAt:    reservation.StartAt.Format("2006-01-02 15:04:05"),
		EndAt:     reservation.EndAt.Format("2006-01-02 15:04:05"),
		Purpose:   reservation.Purpose,
		Status:    reservation.Status,
		ApproveId: halpers.ConvertToInt64(reservation.ApproveId),
		File:      path.Join(config.Cfg.BaseUrl, reservation.File),
		CreatedAt: reservation.CreatedAt,
		UpdatedAt: reservation.UpdatedAt,
	}
}

func (r RoomReservationServiceImpl) Delete(ctx context.Context, id int64) {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	_, err = r.RoomReservationRepository.FindById(ctx, tx, id)
	halpers.IfPanicError(err)

	r.RoomReservationRepository.Delete(ctx, tx, id)
}
