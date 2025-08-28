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

type RoomControllerImpl struct {
	services.RoomService
}

func NewRoomControllerImpl(roomService services.RoomService) *RoomControllerImpl {
	return &RoomControllerImpl{RoomService: roomService}
}

func (r RoomControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var filter web.RoomFilter

	if search := ctx.Query("search"); search != "" {
		filter.Search = &search
	}

	if roomCategoryId := ctx.Query("roomCategoryId"); roomCategoryId != "" {
		filter.RoomCategoryId = &roomCategoryId
	}

	if condition := ctx.Query("condition"); condition != "" {
		filter.Condition = &condition
	}

	rooms := r.RoomService.FindAll(ctx.Context(), filter)

	if rooms == nil {
		rooms = make([]web.RoomResponse, 0)
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   rooms,
	}

	return ctx.JSON(webResponse)
}

func (r RoomControllerImpl) FindById(ctx *fiber.Ctx) error {
	roomId := ctx.Params("roomId")
	id, err := strconv.Atoi(roomId)
	halpers.IfPanicError(err)

	room := r.RoomService.FindById(ctx.Context(), int64(id))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   room,
	}
	return ctx.JSON(webResponse)
}

func (r RoomControllerImpl) Create(ctx *fiber.Ctx) error {
	var roomCreateRequest request.RoomCreateRequest

	err := ctx.BodyParser(&roomCreateRequest)
	halpers.IfPanicError(err)

	form, err := ctx.MultipartForm()
	halpers.IfPanicError(err)

	roomCreateRequest.Images = form.File["images"]
	room := r.RoomService.Create(ctx.Context(), roomCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   room,
	}
	return ctx.JSON(webResponse)
}

func (r RoomControllerImpl) Update(ctx *fiber.Ctx) error {
	var roomUpdateRequest request.RoomUpdateRequest

	err := ctx.BodyParser(&roomUpdateRequest)
	halpers.IfPanicError(err)

	roomId := ctx.Params("roomId")
	id, err := strconv.Atoi(roomId)
	halpers.IfPanicError(err)

	roomUpdateRequest.Id = int64(id)

	room := r.RoomService.Update(ctx.Context(), roomUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   room,
	}
	return ctx.JSON(webResponse)
}

func (r RoomControllerImpl) Delete(ctx *fiber.Ctx) error {
	roomId := ctx.Params("roomId")
	id, err := strconv.Atoi(roomId)
	halpers.IfPanicError(err)

	r.RoomService.Delete(ctx.Context(), int64(id))

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}
	return ctx.JSON(webResponse)
}
