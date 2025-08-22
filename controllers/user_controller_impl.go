package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"

	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
	"github.com/alfianyulianto/go-room-managament/services"
)

type UserControllerImpl struct {
	services.UserService
}

func NewUserControllerImpl(userService services.UserService) *UserControllerImpl {
	return &UserControllerImpl{UserService: userService}
}

func (u UserControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var filter web.UserFilter

	if search := ctx.Query("search"); search != "" {
		filter.Search = &search
	}

	users := u.UserService.FindAll(ctx.Context(), filter)

	if users == nil {
		users = make([]web.UserResponse, 0)
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   users,
	}

	return ctx.JSON(webResponse)
}

func (u UserControllerImpl) FindById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	halpers.IfPanicError(err)

	user := u.UserService.FindById(ctx.Context(), int64(id))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   user,
	}
	return ctx.JSON(webResponse)
}

func (u UserControllerImpl) Create(ctx *fiber.Ctx) error {
	var userCreateRequest request.UserCreateRequest

	err := ctx.BodyParser(&userCreateRequest)
	halpers.IfPanicError(err)

	user := u.UserService.Create(ctx.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   user,
	}
	return ctx.JSON(webResponse)
}

func (u UserControllerImpl) Update(ctx *fiber.Ctx) error {
	var userUpdateRequest request.UserUpdateRequest
	err := ctx.BodyParser(&userUpdateRequest)
	halpers.IfPanicError(err)

	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	halpers.IfPanicError(err)
	userUpdateRequest.Id = int64(id)

	user := u.UserService.Update(ctx.Context(), userUpdateRequest)
	WebResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   user,
	}
	return ctx.JSON(WebResponse)
}

func (u UserControllerImpl) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	halpers.IfPanicError(err)

	u.UserService.Delete(ctx.Context(), int64(id))

	WebResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}
	return ctx.JSON(WebResponse)
}
