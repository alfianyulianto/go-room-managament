package router

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/alfianyulianto/go-room-managament/controllers"
	"github.com/alfianyulianto/go-room-managament/exception"
	"github.com/alfianyulianto/go-room-managament/middleware"
	"github.com/alfianyulianto/go-room-managament/util"
)

type RouterConfig struct {
	controllers.AuthController
	controllers.UserController
	controllers.RoomCategoryController
	controllers.RoomController
	controllers.RoomImageController
	controllers.RoomReservationController
}

func NewRouter(config RouterConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		BodyLimit:    20 * 1024 * 1024,
	})

	app.Use(recover2.New())

	app.Static("/uploads", "./uploads")

	tokenUtil := util.NewTokenUtil()
	authMiddleware := middleware.NewAuth(tokenUtil)

	AuthRoute(app.Group("/auth"), config.AuthController)

	app.Use(authMiddleware)
	UserRoute(app.Group("/users"), config.UserController)
	RoomCategoryRouter(app.Group("/room-categories"), config.RoomCategoryController)
	RoomRoute(app.Group("/rooms"), config.RoomController)
	RoomImageRoute(app.Group("/rooms/:roomId/images"), config.RoomImageController)
	RoomReservationRoute(app.Group("/room-reservations"), config.RoomReservationController)

	return app
}
