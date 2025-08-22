package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"

	"github.com/alfianyulianto/go-room-managament/model/web"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if _, ok := err.(validator.ValidationErrors); ok {
		return validatorErrors(ctx, err.(validator.ValidationErrors))
	} else if _, ok = err.(*NotFoundError); ok {
		return notFoundError(ctx, err.(*NotFoundError))
	} else {
		return internalServerError(ctx, err)
	}
}

func notFoundError(ctx *fiber.Ctx, err *NotFoundError) error {
	ctx.Status(http.StatusNotFound)

	webResponse := web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   err.Message,
	}
	return ctx.JSON(webResponse)
}

func validatorErrors(ctx *fiber.Ctx, errors validator.ValidationErrors) error {
	ctx.Status(http.StatusBadRequest)

	var validationErrors []web.ValidationErrorResponse
	for _, fieldError := range errors {
		validationErrors = append(validationErrors, web.ValidationErrorResponse{
			Field:   fieldError.Field(),
			Message: fieldError.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Data:   validationErrors,
	}
	return ctx.JSON(webResponse)
}

func internalServerError(ctx *fiber.Ctx, err error) error {
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Data:   err.Error(),
	}
	return ctx.Status(http.StatusInternalServerError).JSON(webResponse)
}
