package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"

	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/web"
	"github.com/alfianyulianto/go-room-managament/util"
)

func NewAuth(tokenUtil *util.TokenUtil) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenRequest := ctx.Get("Authorization")

		if !strings.HasPrefix(tokenRequest, "Bearer ") {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
			}

			return ctx.JSON(webResponse)
		}

		auth, err := tokenUtil.ParseToken(strings.TrimPrefix(tokenRequest, "Bearer "))
		if err != nil {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: err.Error(),
			}

			return ctx.JSON(webResponse)
		}

		ctx.Locals("auth", auth)

		ctx.SetUserContext(context.WithValue(ctx.UserContext(), "auth", auth))

		return ctx.Next()

	}
}

func GetAuth(ctx *fiber.Ctx) *domain.Auth {
	return ctx.Locals("auth").(*domain.Auth)
}
