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
	controllers.RoomReservationController
}

func MethodOverride() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// cek dulu content type
		if c.Is("multipart/form-data") || c.Is("application/x-www-form-urlencoded") {
			if m := c.FormValue("_method"); m != "" {
				c.Request().Header.SetMethod(m)
			}
		} else if m := c.Query("_method"); m != "" {
			c.Request().Header.SetMethod(m)
		} else if m := c.Get("X-HTTP-Method-Override"); m != "" {
			c.Request().Header.SetMethod(m)
		}
		return c.Next()
	}
}

func NewRouter(config RouterConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		BodyLimit:    20 * 1024 * 1024,
	})

	app.Use(recover2.New())

	app.Use(MethodOverride())

	app.Static("/uploads", "./uploads")

	UserRoute(app.Group("/users"), config.UserController)
	RoomCategoryRouter(app.Group("/room-categories"), config.RoomCategoryController)
	RoomRoute(app.Group("/rooms"), config.RoomController)
	RoomImageRoute(app.Group("/rooms/:roomId/images"), config.RoomImageController)
	RoomReservationRoute(app.Group("/room-reservations"), config.RoomReservationController)

	return app
}
