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

	"github.com/alfianyulianto/go-room-managament/config"
	"github.com/alfianyulianto/go-room-managament/exception"
	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
	"github.com/alfianyulianto/go-room-managament/repositories"
	"github.com/alfianyulianto/go-room-managament/storage"
)

type RoomServiceImpl struct {
	repositories.RoomRepository
	repositories.RoomCategoryRepository
	repositories.RoomImageRepository
	storage.FileStorage
	*sql.DB
	*validator.Validate
}

func NewRoomServiceImpl(RoomRepository repositories.RoomRepository, RoomCategoryRepository repositories.RoomCategoryRepository, RoomImageRepository repositories.RoomImageRepository, fileStorage storage.FileStorage, DB *sql.DB, validate *validator.Validate) *RoomServiceImpl {
	err := validate.RegisterValidation("uniqueRoomNameCreate", func(field validator.FieldLevel) bool {
		name := field.Parent().FieldByName("Name").String()
		room, _ := RoomRepository.FindByName(context.Background(), DB, name)
		if room != nil {
			return false
		}
		return true
	})

	err = validate.RegisterValidation("uniqueRoomCodeCreate", func(field validator.FieldLevel) bool {
		code := field.Parent().FieldByName("Code").String()
		room, _ := RoomRepository.FindByCode(context.Background(), DB, code)
		if room != nil {
			return false
		}
		return true
	})

	err = validate.RegisterValidation("uniqueRoomNameUpdate", func(field validator.FieldLevel) bool {
		id := field.Parent().FieldByName("Id").Int()
		name := field.Field().String()
		room, _ := RoomRepository.FindByName(context.Background(), DB, name)
		if room != nil {
			if room.Id != id {
				return false
			}
			return true
		}
		return true
	})

	err = validate.RegisterValidation("uniqueRoomCodeUpdate", func(field validator.FieldLevel) bool {
		id := field.Parent().FieldByName("Id").Int()
		code := field.Field().String()
		room, _ := RoomRepository.FindByCode(context.Background(), DB, code)
		if room != nil {
			if room.Id != id {
				return false
			}
			return true
		}
		return true
	})

	err = validate.RegisterValidation("existRoomCategory", func(field validator.FieldLevel) bool {
		roomCategoryId := field.Parent().FieldByName("RoomCategoryId").Int()
		roomCategory, _ := RoomCategoryRepository.FindById(context.Background(), DB, roomCategoryId)
		if roomCategory != nil {
			return true
		}
		return false
	})

	err = validate.RegisterValidation("files", func(fl validator.FieldLevel) bool {
		files, ok := fl.Field().Interface().([]*multipart.FileHeader)
		if !ok {
			return false
		}

		for _, file := range files {
			if file == nil {
				return false
			}

			fmt.Println("cek file:", file.Filename, "size:", file.Size)

			if file.Size > 5<<20 {
				return false
			}

			ext := strings.ToLower(filepath.Ext(file.Filename))
			if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
				return false
			}
		}

		return true
	})

	halpers.IfPanicError(err)

	return &RoomServiceImpl{RoomRepository: RoomRepository, RoomCategoryRepository: RoomCategoryRepository, RoomImageRepository: RoomImageRepository, FileStorage: fileStorage, DB: DB, Validate: validate}
}

func (r RoomServiceImpl) FindAll(ctx context.Context, filter web.RoomFilter) []web.RoomResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer tx.Rollback()

	rooms := r.RoomRepository.FindAll(ctx, tx, filter)
	var roomResponses []web.RoomResponse
	for _, room := range rooms {
		roomResponses = append(roomResponses, web.RoomResponse{
			Id:               room.Id,
			RoomCategoryId:   room.RoomCategoryId,
			RoomCategoryName: room.RoomCategoryName,
			Code:             room.Code,
			Name:             room.Name,
			Condition:        room.Condition,
			Note:             room.Note,
			CreatedAt:        room.CreatedAt,
			UpdatedAt:        room.UpdatedAt,
		})
	}
	return roomResponses
}

func (r RoomServiceImpl) FindById(ctx context.Context, id int64) web.RoomResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer tx.Rollback()

	room, err := r.RoomRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(&exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return web.RoomResponse{
		Id:               room.Id,
		RoomCategoryId:   room.RoomCategoryId,
		RoomCategoryName: room.RoomCategoryName,
		Code:             room.Code,
		Name:             room.Name,
		Condition:        room.Condition,
		Note:             room.Note,
		CreatedAt:        room.CreatedAt,
		UpdatedAt:        room.UpdatedAt,
	}

}

func (r RoomServiceImpl) Create(ctx context.Context, request request.RoomCreateRequest) web.RoomResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	err = r.Validate.Struct(request)
	halpers.IfPanicError(err)

	room := r.RoomRepository.Save(ctx, tx, domain.Room{
		RoomCategoryId: request.RoomCategoryId,
		Code:           request.Code,
		Name:           request.Name,
		Condition:      request.Condition,
		Note:           request.Note,
	})

	var roomImages []web.RoomImageResponse
	for _, image := range request.Images {
		filePath, err := r.FileStorage.SaveFile(image, "/room_images")
		halpers.IfPanicError(err)
		roomImage := r.RoomImageRepository.Save(ctx, tx, domain.RoomImage{
			RoomId: room.Id,
			Path:   filePath,
		})

		roomImages = append(roomImages, web.RoomImageResponse{
			Id:        roomImage.Id,
			RoomId:    roomImage.RoomId,
			Path:      path.Join(config.Cfg.BaseUrl, filePath),
			CreatedAt: roomImage.CreatedAt,
			UpdatedAt: roomImage.UpdatedAt,
		})
	}

	return web.RoomResponse{
		Id:               room.Id,
		RoomCategoryId:   room.RoomCategoryId,
		RoomCategoryName: room.RoomCategoryName,
		Code:             room.Code,
		Name:             room.Name,
		Condition:        room.Condition,
		Note:             room.Note,
		CreatedAt:        room.CreatedAt,
		UpdatedAt:        room.UpdatedAt,
		Images:           roomImages,
	}
}

func (r RoomServiceImpl) Update(ctx context.Context, request request.RoomUpdateRequest) web.RoomResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer tx.Rollback()

	err = r.Validate.Struct(request)
	halpers.IfPanicError(err)

	// check room by id
	_, err = r.RoomRepository.FindById(ctx, tx, request.Id)
	halpers.IfPanicError(err)

	room := r.RoomRepository.Update(ctx, tx, domain.Room{
		Id:             request.Id,
		RoomCategoryId: request.RoomCategoryId,
		Code:           request.Code,
		Name:           request.Name,
		Condition:      request.Condition,
		Note:           request.Note,
	})

	return web.RoomResponse{
		Id:               room.Id,
		RoomCategoryId:   room.RoomCategoryId,
		RoomCategoryName: room.RoomCategoryName,
		Code:             room.Code,
		Name:             room.Name,
		Condition:        room.Condition,
		Note:             room.Note,
		CreatedAt:        room.CreatedAt,
		UpdatedAt:        room.UpdatedAt,
	}
}

func (r RoomServiceImpl) Delete(ctx context.Context, id int64) {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	// check room by id
	_, err = r.RoomRepository.FindById(ctx, tx, id)
	halpers.IfPanicError(err)

	r.RoomRepository.Delete(ctx, tx, id)

}
