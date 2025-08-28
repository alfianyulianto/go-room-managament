package exception

import (
	"fmt"
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
		message := fieldError.Error()
		// Custom error messages per tag
		switch fieldError.Tag() {
		case "uniqueUserNameCreate", "uniqueUserEmailCreate", "uniqueRoomNameCreate", "uniqueRoomCodeCreate", "uniqueRoomNameUpdate", "uniqueRoomCodeUpdate":
			message = fmt.Sprintf("%s has alredy been taken", fieldError.Field())
		case "existRoomCategory":
			message = "room category does not exist"
		case "required":
			message = fmt.Sprintf("%s is required", fieldError.Field())
		case "email":
			message = fmt.Sprintf("%s must be a valid email", fieldError.Field())
		case "files", "file":
			message = fmt.Sprintf("%s may not be greater than 5MB", "Image")
		case "eq=Admin|eq=User":
			message = fmt.Sprintf("selected %s is invalid", fieldError.Field())
		case "eq=Baik|eq=Rusak Ringan|eq=Rusak Sedang|eq=Rusak Berat":
			message = fmt.Sprintf("selected %s is invalid", fieldError.Field())

		}

		validationErrors = append(validationErrors, web.ValidationErrorResponse{
			Field:   fieldError.Field(),
			Message: message,
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
