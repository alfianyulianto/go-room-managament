package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alfianyulianto/go-room-managament/controllers"
)

func UserRoute(router fiber.Router, controller controllers.UserController) {
	router.Get("/", controller.FindAll)
	router.Get("/:userId", controller.FindById)
	router.Post("/", controller.Create)
	router.Put("/:userId", controller.Update)
	router.Delete("/:userId", controller.Delete)
}
