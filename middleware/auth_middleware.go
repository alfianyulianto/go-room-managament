package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"

	"github.com/alfianyulianto/go-room-managament/model/domain"
	"github.com/alfianyulianto/go-room-managament/model/web"
	"github.com/alfianyulianto/go-room-managament/util"
)

func NewAuth(tokenUtil *util.TokenUtil) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenRequest := ctx.Get("Authorization")
		fmt.Println(tokenRequest)

		if tokenRequest == "" {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
			}

			return ctx.JSON(webResponse)
		}

		if len(tokenRequest) > 7 && tokenRequest[:7] == "Bearer " {
			tokenRequest = tokenRequest[7:]
		}

		auth, err := tokenUtil.ParseToken(tokenRequest)
		if err != nil {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: err.Error(),
			}

			return ctx.JSON(webResponse)
		}

		ctx.Locals("auth", auth)

		return ctx.Next()

	}
}

func GetAuth(ctx *fiber.Ctx) *domain.Auth {
	return ctx.Locals("auth").(*domain.Auth)
}
