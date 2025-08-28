package router

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/alfianyulianto/go-room-managament/controllers"
	"github.com/alfianyulianto/go-room-managament/exception"
)

type RouterConfig struct {
	controllers.UserController
	controllers.RoomCategoryController
	controllers.RoomController
	controllers.RoomImageController
}

func NewRouter(config RouterConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		BodyLimit:    20 * 1024 * 1024,
	})

	app.Use(recover2.New())

	app.Static("/uploads", "./uploads")

	UserRoute(app.Group("/users"), config.UserController)
	RoomCategoryRouter(app.Group("/room-categories"), config.RoomCategoryController)
	RoomRoute(app.Group("/rooms"), config.RoomController)
	RoomImageRoute(app.Group("/rooms/:roomId/images"), config.RoomImageController)

	return app
}
