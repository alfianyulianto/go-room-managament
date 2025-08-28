package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alfianyulianto/go-room-managament/controllers"
)

func RoomImageRoute(router fiber.Router, controller controllers.RoomImageController) {
	router.Get("/", controller.FindAll)
	router.Get("/:roomImageId", controller.FindById)
	router.Post("/", controller.Create)
	router.Delete("/:roomImageId", controller.Delete)
}
