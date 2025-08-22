package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alfianyulianto/go-room-managament/controllers"
)

func RoomCategoryRouter(router fiber.Router, controller controllers.RoomCategoryController) {
	router.Get("/", controller.FindAll)
	router.Get("/:roomCategoryId", controller.FindById)
	router.Post("/", controller.Create)
	router.Put("/:roomCategoryId", controller.Update)
	router.Delete("/:roomCategoryId", controller.Delete)
}
