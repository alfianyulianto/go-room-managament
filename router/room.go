package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alfianyulianto/go-room-managament/controllers"
)

func RoomRoute(router fiber.Router, controller controllers.RoomController) {
	router.Get("/", controller.FindAll)
	router.Get("/:roomId", controller.FindById)
	router.Post("/", controller.Create)
	router.Put("/:roomId", controller.Update)
	router.Delete("/:roomId", controller.Delete)
}
