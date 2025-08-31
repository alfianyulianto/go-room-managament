package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"

	"github.com/alfianyulianto/go-room-managament/exception"
	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
	"github.com/alfianyulianto/go-room-managament/repositories"
	"github.com/alfianyulianto/go-room-managament/util"
)

type AuthServiceImpl struct {
	repositories.UserRepository
	*sql.DB
	*validator.Validate
	*util.TokenUtil
}

func NewAuthServiceImpl(userRepository repositories.UserRepository, DB *sql.DB, validate *validator.Validate, tokenUtil *util.TokenUtil) *AuthServiceImpl {
	err := validate.RegisterValidation("uniqueUserNameCreate", func(field validator.FieldLevel) bool {
		name := field.Parent().FieldByName("Name").String()
		user, _ := userRepository.FindByName(context.Background(), DB, name)
		if user != nil {
			return false
		}
		return true
	})

	err = validate.RegisterValidation("uniqueUserEmailCreate", func(field validator.FieldLevel) bool {
		email := field.Parent().FieldByName("Email").String()
		user, _ := userRepository.FindByEmail(context.Background(), DB, email)
		if user != nil {
			return false
		}
		return true
	})

	halpers.IfPanicError(err)

	return &AuthServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate, TokenUtil: tokenUtil}
}

func (u AuthServiceImpl) Login(ctx context.Context, request request.LoginRequest) (*web.TokenResponse, error) {
	tx, err := u.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	err = u.Validate.Struct(request)
	halpers.IfPanicError(err)

	user, err := u.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(&exception.NotFoundError{Message: err.Error()})
	}

	if user.Password == request.Password {
		return nil, errors.New("password is invalid")
	}

	token, err := u.TokenUtil.CreateToken(domain.Auth{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Level: user.Level,
	})
	halpers.IfPanicError(err)

	return &web.TokenResponse{
		Token: token,
	}, nil
}

func (u AuthServiceImpl) Register(ctx context.Context, request request.RegisterRequest) web.UserResponse {
	tx, err := u.DB.Begin()
	halpers.IfPanicError(err)
	defer halpers.CommitOrRollback(tx)

	fmt.Println(request)

	err = u.Validate.Struct(request)
	halpers.IfPanicError(err)

	user := u.UserRepository.Save(ctx, tx, domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: halpers.HasPassword(request.Password),
		Level:    "User",
	})

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
