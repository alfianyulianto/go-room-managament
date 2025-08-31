package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alfianyulianto/go-room-managament/controllers"
)

func RoomReservationRoute(router fiber.Router, controller controllers.RoomReservationController) {
	router.Get("/", controller.FindAll)
	router.Get("/:roomReservationId", controller.FindById)
	router.Post("/", controller.Create)
	router.Put("/:roomReservationId", controller.Update)
	router.Delete("/:roomReservationId", controller.Delete)
	router.Post("/:roomReservationId/set-accepted", controller.Accepted)
	router.Post("/:roomReservationId/set-rejected", controller.Rejected)
}
