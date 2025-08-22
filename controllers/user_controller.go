package controllers

import "github.com/gofiber/fiber/v2"

type UserController interface {
	FindAll(*fiber.Ctx) error
	FindById(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
}
