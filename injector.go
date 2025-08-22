//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"

	"github.com/alfianyulianto/go-room-managament/app"
	"github.com/alfianyulianto/go-room-managament/controllers"
	"github.com/alfianyulianto/go-room-managament/repositories"
	"github.com/alfianyulianto/go-room-managament/router"
	"github.com/alfianyulianto/go-room-managament/services"
)

var roomCategorySet = wire.NewSet(
	repositories.NewRoomCategoryRepositoryImp,
	wire.Bind(new(repositories.RoomCategoryRepository), new(*repositories.RoomCategoryRepositoryImpl)),

	services.NewRoomCategoryServiceImpl,
	wire.Bind(new(services.RoomCategoryService), new(*services.RoomCategoryServiceImpl)),

	controllers.NewRoomCategoryControllerImpl,
	wire.Bind(new(controllers.RoomCategoryController), new(*controllers.RoomCategoryControllerImpl)),
)

var userSet = wire.NewSet(
	repositories.NewUserRepositoryImpl,
	wire.Bind(new(repositories.UserRepository), new(*repositories.UserRepositoryImpl)),

	services.NewUserServiceImpl,
	wire.Bind(new(services.UserService), new(*services.UserServiceImpl)),

	controllers.NewUserControllerImpl,
	wire.Bind(new(controllers.UserController), new(*controllers.UserControllerImpl)),
)

func NewInitializedServer(options ...validator.Option) *fiber.App {
	wire.Build(
		app.NewDB,
		validator.New,
		roomCategorySet,
		userSet,
		wire.Struct(new(router.RouterConfig), "*"),
		router.NewRouter,
	)
	return nil
}
