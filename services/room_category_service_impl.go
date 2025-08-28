package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"

	"github.com/alfianyulianto/go-room-managament/exception"
	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
	"github.com/alfianyulianto/go-room-managament/repositories"
)

type RoomCategoryServiceImpl struct {
	RoomCategoryRepository repositories.RoomCategoryRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewRoomCategoryServiceImpl(roomCategoryRepository repositories.RoomCategoryRepository, DB *sql.DB, validate *validator.Validate) *RoomCategoryServiceImpl {
	//validation
	err := validate.RegisterValidation("uniqueRoomCategoryCreate", func(field validator.FieldLevel) bool {
		name := field.Field().String()
		roomCategory, _ := roomCategoryRepository.FindByName(context.Background(), DB, name)
		if roomCategory != nil {
			return false
		}
		return true
	})

	err = validate.RegisterValidation("uniqueRoomCategoryUpdate", func(field validator.FieldLevel) bool {
		id := field.Parent().FieldByName("Id").Int()
		name := field.Field().String()
		roomCategory, _ := roomCategoryRepository.FindByName(context.Background(), DB, name)
		if roomCategory != nil {
			if roomCategory.Id != id {
				return false
			}
			return true
		}
		return true
	})

	halpers.IfPanicError(err)

	return &RoomCategoryServiceImpl{RoomCategoryRepository: roomCategoryRepository, DB: DB, Validate: validate}
}

func (r *RoomCategoryServiceImpl) FindAll(ctx context.Context, filter web.RoomCategoryFilter) []web.RoomCategoryResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	roomCategories := r.RoomCategoryRepository.FindAll(ctx, tx, filter)
	var roomCategoriesResponses []web.RoomCategoryResponse
	for _, roomCategory := range roomCategories {
		roomCategoriesResponses = append(roomCategoriesResponses, web.RoomCategoryResponse{roomCategory.Id, roomCategory.Name, roomCategory.CreatedAt, roomCategory.UpdatedAt})
	}
	return roomCategoriesResponses
}

func (r *RoomCategoryServiceImpl) FindById(ctx context.Context, id int64) web.RoomCategoryResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	roomCategory, err := r.RoomCategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(&exception.NotFoundError{Message: err.Error()})
	}

	return web.RoomCategoryResponse{
		Id:        roomCategory.Id,
		Name:      roomCategory.Name,
		CreatedAt: roomCategory.CreatedAt,
		UpdatedAt: roomCategory.UpdatedAt,
	}
}

func (r *RoomCategoryServiceImpl) Create(ctx context.Context, request request.RoomCategoryCreateRequest) web.RoomCategoryResponse {
	tx, err := r.DB.Begin()
	defer halpers.CommitOrRollback(tx)

	err = r.Validate.Struct(request)
	halpers.IfPanicError(err)

	roomCategory := r.RoomCategoryRepository.Save(ctx, tx, domain.RoomCategory{
		Name: request.Name,
	})

	return web.RoomCategoryResponse{
		Id:        roomCategory.Id,
		Name:      roomCategory.Name,
		CreatedAt: roomCategory.CreatedAt,
		UpdatedAt: roomCategory.UpdatedAt,
	}
}

func (r *RoomCategoryServiceImpl) Update(ctx context.Context, request request.RoomCategoryUpdateRequest) web.RoomCategoryResponse {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	//validation
	err = r.Validate.Struct(request)
	halpers.IfPanicError(err)

	//check room category by id
	_, err = r.RoomCategoryRepository.FindById(ctx, tx, request.Id)
	halpers.IfPanicError(err)

	roomCategory := r.RoomCategoryRepository.Update(ctx, tx, domain.RoomCategory{
		Id:   request.Id,
		Name: request.Name,
	})

	return web.RoomCategoryResponse{
		Id:        roomCategory.Id,
		Name:      roomCategory.Name,
		CreatedAt: roomCategory.CreatedAt,
		UpdatedAt: roomCategory.UpdatedAt,
	}
}

func (r *RoomCategoryServiceImpl) Delete(ctx context.Context, id int64) {
	tx, err := r.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	// check room category by id
	roomCategory, err := r.RoomCategoryRepository.FindById(ctx, tx, id)
	halpers.IfPanicError(err)

	r.RoomCategoryRepository.Delete(ctx, tx, roomCategory.Id)
}
