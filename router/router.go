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
}

func NewRouter(config RouterConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	app.Use(recover2.New())

	UserRoute(app.Group("/users"), config.UserController)
	RoomCategoryRouter(app.Group("/room-categories"), config.RoomCategoryController)

	return app
}
