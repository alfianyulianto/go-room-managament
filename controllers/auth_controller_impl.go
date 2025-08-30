package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"

	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
	"github.com/alfianyulianto/go-room-managament/services"
)

type AuthControllerImpl struct {
	services.AuthService
}

func NewAuthControllerImpl(authService services.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{AuthService: authService}
}

func (a *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	var loginRequest request.LoginRequest
	err := ctx.BodyParser(&loginRequest)
	halpers.IfPanicError(err)

	token, err := a.AuthService.Login(ctx.Context(), loginRequest)

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: http.StatusText(fiber.StatusOK),
		Data:   token,
	}

	return ctx.JSON(webResponse)
}

func (a *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	var registerRequest request.RegisterRequest
	err := ctx.BodyParser(&registerRequest)
	halpers.IfPanicError(err)

	user := a.AuthService.Register(ctx.Context(), registerRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   user,
	}

	return ctx.JSON(webResponse)
}
