package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"

	"github.com/alfianyulianto/go-room-managament/halpers"
	"github.com/alfianyulianto/go-room-managament/model/request"
	"github.com/alfianyulianto/go-room-managament/model/web"
	"github.com/alfianyulianto/go-room-managament/services"
)

type RoomCategoryControllerImpl struct {
	RoomCategoryService services.RoomCategoryService
}

func NewRoomCategoryControllerImpl(roomCategoryService services.RoomCategoryService) *RoomCategoryControllerImpl {
	return &RoomCategoryControllerImpl{RoomCategoryService: roomCategoryService}
}

func (r RoomCategoryControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var filter web.RoomCategoryFilter

	if search := ctx.Query("search"); search != "" {
		filter.Search = &search
	}
	roomCategories := r.RoomCategoryService.FindAll(ctx.Context(), filter)
	if roomCategories == nil {
		roomCategories = make([]web.RoomCategoryResponse, 0)
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomCategories,
	}

	err := ctx.JSON(webResponse)
	return err
}

func (r RoomCategoryControllerImpl) FindById(ctx *fiber.Ctx) error {
	roomCategoryId := ctx.Params("roomCategoryId")
	id, err2 := strconv.Atoi(roomCategoryId)
	halpers.IfPanicError(err2)
	roomCategory := r.RoomCategoryService.FindById(ctx.Context(), int64(id))

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomCategory,
	}
	err := ctx.JSON(webResponse)
	return err
}

func (r RoomCategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	var roomCategoryCreateRequest request.RoomCategoryCreateRequest

	err := ctx.BodyParser(&roomCategoryCreateRequest)
	halpers.IfPanicError(err)

	roomCategory := r.RoomCategoryService.Create(ctx.Context(), roomCategoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomCategory,
	}
	return ctx.JSON(webResponse)
}

func (r RoomCategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	var roomCategoryUpdateRequest request.RoomCategoryUpdateRequest

	err := ctx.BodyParser(&roomCategoryUpdateRequest)
	halpers.IfPanicError(err)

	roomCategoryId := ctx.Params("roomCategoryId")
	id, err2 := strconv.Atoi(roomCategoryId)
	halpers.IfPanicError(err2)

	roomCategoryUpdateRequest.Id = int64(id)

	roomCategory := r.RoomCategoryService.Update(ctx.Context(), roomCategoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomCategory,
	}
	return ctx.JSON(webResponse)
}

func (r RoomCategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	roomCategoryId := ctx.Params("roomCategoryId")
	id, err2 := strconv.Atoi(roomCategoryId)
	halpers.IfPanicError(err2)
	r.RoomCategoryService.Delete(ctx.Context(), int64(id))

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   []string{},
	}
	err := ctx.JSON(webResponse)
	return err
}
