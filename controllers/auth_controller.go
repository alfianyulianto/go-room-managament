package controllers

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(*fiber.Ctx) error
}
