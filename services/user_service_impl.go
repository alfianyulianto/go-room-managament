package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"

	"github.com/alfianyulianto/go-room-managament/exception"
	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
	"github.com/alfianyulianto/go-room-managament/repositories"
)

type UserServiceImpl struct {
	repositories.UserRepository
	*sql.DB
	*validator.Validate
}

func NewUserServiceImpl(UserRepository repositories.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {

	err := validate.RegisterValidation("uniqueUserNameCreate", func(field validator.FieldLevel) bool {
		name := field.Parent().FieldByName("Name").String()
		user, _ := UserRepository.FindByName(context.Background(), DB, name)
		if user != nil {
			return false
		}
		return true
	})

	err = validate.RegisterValidation("uniqueUserEmailCreate", func(field validator.FieldLevel) bool {
		email := field.Parent().FieldByName("Email").String()
		user, _ := UserRepository.FindByEmail(context.Background(), DB, email)
		if user != nil {
			return false
		}
		return true
	})

	err = validate.RegisterValidation("uniqueUserNameUpdate", func(field validator.FieldLevel) bool {
		id := field.Parent().FieldByName("Id").Int()
		name := field.Field().String()
		user, _ := UserRepository.FindByName(context.Background(), DB, name)
		if user != nil {
			if user.Id != id {
				return false
			}
			return true
		}
		return true
	})

	err = validate.RegisterValidation("uniqueUserEmailUpdate", func(field validator.FieldLevel) bool {
		id := field.Parent().FieldByName("Id").Int()
		email := field.Field().String()
		user, _ := UserRepository.FindByEmail(context.Background(), DB, email)
		if user != nil {
			if user.Id != id {
				return false
			}
			return true
		}
		return true
	})
	halpers.IfPanicError(err)

	return &UserServiceImpl{UserRepository: UserRepository, DB: DB, Validate: validate}
}

func (u UserServiceImpl) FindAll(ctx context.Context, filter web.UserFilter) []web.UserResponse {
	tx, err := u.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	users := u.UserRepository.FindAll(ctx, tx, filter)
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, web.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Level:     user.Level,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return userResponses
}

func (u UserServiceImpl) FindById(ctx context.Context, id int64) web.UserResponse {
	tx, err := u.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	user, err := u.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(&exception.NotFoundError{Message: err.Error()})
	}

	return web.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Level:     user.Level,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

}

func (u UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) web.UserResponse {
	tx, err := u.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	err = u.Validate.Struct(request)
	halpers.IfPanicError(err)

	fmt.Println(request)

	user := u.UserRepository.Save(ctx, tx, domain.User{
		Name:  request.Name,
		Email: request.Email,
		Phone: request.Phone,
		Level: request.Level,
	})
	return web.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Level:     user.Level,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) web.UserResponse {
	tx, err := u.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	err = u.Validate.Struct(request)
	halpers.IfPanicError(err)

	user, err := u.UserRepository.FindById(ctx, tx, request.Id)
	halpers.IfPanicError(err)

	user = u.UserRepository.Update(ctx, tx, domain.User{
		Id:    request.Id,
		Name:  request.Name,
		Email: request.Email,
		Phone: request.Phone,
		Level: request.Level,
	})

	return web.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Level:     user.Level,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u UserServiceImpl) Delete(ctx context.Context, id int64) {
	tx, err := u.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	_, err = u.UserRepository.FindById(ctx, tx, id)
	halpers.IfPanicError(err)

	u.UserRepository.Delete(ctx, tx, id)
}
