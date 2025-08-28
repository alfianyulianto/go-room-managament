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

type RoomImageServiceImpl struct {
	RoomRepository repositories.RoomRepository
	repositories.RoomImageRepository
	storage.FileStorage
	*sql.DB
	*validator.Validate
}

func NewRoomImageServiceImpl(RoomRepository repositories.RoomRepository, RoomImageRepository repositories.RoomImageRepository, fileStorage storage.FileStorage, DB *sql.DB, validate *validator.Validate) *RoomImageServiceImpl {
	err := validate.RegisterValidation("file", func(fl validator.FieldLevel) bool {
		file, ok := fl.Field().Interface().(multipart.FileHeader)
		if !ok {
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

		return true
	})
	halpers.IfPanicError(err)

	return &RoomImageServiceImpl{RoomRepository: RoomRepository, RoomImageRepository: RoomImageRepository, FileStorage: fileStorage, DB: DB, Validate: validate}
}

func (r RoomImageServiceImpl) ensureRoomExists(ctx context.Context, tx *sql.Tx, roomId int64) {
	_, err := r.RoomRepository.FindById(ctx, tx, roomId)
	if err != nil {
		panic(&exception.NotFoundError{Message: err.Error()})
	}
}

func (r RoomImageServiceImpl) FindAll(ctx context.Context, roomId int64) []web.RoomImageResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)

	r.ensureRoomExists(ctx, tx, roomId)

	roomImages := r.RoomImageRepository.FindAll(ctx, tx, roomId)
	var roomImageResponses []web.RoomImageResponse
	for _, roomImage := range roomImages {
		roomImageResponses = append(roomImageResponses, web.RoomImageResponse{
			Id:        roomImage.Id,
			RoomId:    roomImage.RoomId,
			Path:      path.Join(config.Cfg.BaseUrl, roomImage.Path),
			CreatedAt: roomImage.CreatedAt,
			UpdatedAt: roomImage.UpdatedAt,
		})
	}
	return roomImageResponses
}

func (r RoomImageServiceImpl) FindById(ctx context.Context, roomId int64, id int64) web.RoomImageResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)

	r.ensureRoomExists(ctx, tx, roomId)

	roomImage, err := r.RoomImageRepository.FindById(ctx, tx, roomId, id)
	if err != nil {
		panic(&exception.NotFoundError{Message: err.Error()})
	}
	return web.RoomImageResponse{
		Id:        roomImage.Id,
		RoomId:    roomImage.RoomId,
		Path:      path.Join(config.Cfg.BaseUrl, roomImage.Path),
		CreatedAt: roomImage.CreatedAt,
	}
}

func (r RoomImageServiceImpl) Create(ctx context.Context, request request.RoomImageCreateRequest) web.RoomImageResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	r.ensureRoomExists(ctx, tx, request.RoomId)

	err = r.Validate.Struct(request)
	halpers.IfPanicError(err)

	pathSaved, err := r.FileStorage.SaveFile(request.Image, "room_images")
	halpers.IfPanicError(err)

	roomImage := r.RoomImageRepository.Save(ctx, tx, domain.RoomImage{
		RoomId: request.RoomId,
		Path:   pathSaved,
	})
	return web.RoomImageResponse{
		Id:        roomImage.Id,
		RoomId:    roomImage.RoomId,
		Path:      path.Join(config.Cfg.BaseUrl, pathSaved),
		CreatedAt: roomImage.CreatedAt,
		UpdatedAt: roomImage.UpdatedAt,
	}
}

func (r RoomImageServiceImpl) Update(ctx context.Context, request request.RoomImageUpdateRequest) web.RoomImageResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	r.ensureRoomExists(ctx, tx, request.RoomId)

	err = r.Validate.Struct(request)
	halpers.IfPanicError(err)

	pathSaved, err := r.FileStorage.SaveFile(request.Image, "room_images")
	halpers.IfPanicError(err)
	roomImage := r.RoomImageRepository.Update(ctx, tx, domain.RoomImage{
		Id:   request.Id,
		Path: pathSaved,
	})

	return web.RoomImageResponse{
		Id:        roomImage.Id,
		RoomId:    roomImage.RoomId,
		Path:      path.Join(config.Cfg.BaseUrl, pathSaved),
		CreatedAt: roomImage.CreatedAt,
	}
}

func (r RoomImageServiceImpl) Delete(ctx context.Context, roomId int64, id int64) {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	r.ensureRoomExists(ctx, tx, roomId)

	_, err = r.RoomImageRepository.FindById(ctx, tx, roomId, id)
	if err != nil {
		panic(&exception.NotFoundError{Message: err.Error()})
	}

	r.RoomImageRepository.Delete(ctx, tx, roomId, id)
}
