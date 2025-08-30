package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alfianyulianto/go-room-managament/controllers"
)

func AuthRoute(router fiber.Router, controller controllers.AuthController) {
	router.Post("/login", controller.Login)
	router.Post("/register", controller.Register)
}
