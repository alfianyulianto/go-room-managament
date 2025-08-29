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

type RoomReservationControllerImpl struct {
	services.RoomReservationService
}

func NewRoomReservationControllerImpl(roomReservationService services.RoomReservationService) *RoomReservationControllerImpl {
	return &RoomReservationControllerImpl{RoomReservationService: roomReservationService}
}

func (r RoomReservationControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var filter web.RoomReservationFilter

	if search := ctx.Query("search"); search != "" {
		filter.Search = &search
	}

	if startAt := ctx.Query("startAt"); startAt != "" {
		filter.StartAt = &startAt
	}

	if endAt := ctx.Query("endAt"); endAt != "" {
		filter.EndAt = &endAt
	}

	if status := ctx.Query("status"); status != "" {
		filter.Status = &status
	}

	roomReservations := r.RoomReservationService.FindAll(ctx.Context(), filter)

	if roomReservations == nil {
		roomReservations = make([]web.RoomReservationResponse, 0)
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomReservations,
	}
	return ctx.JSON(webResponse)
}

func (r RoomReservationControllerImpl) FindById(ctx *fiber.Ctx) error {
	roomReservationId := ctx.Params("roomReservationId")
	id, err := strconv.Atoi(roomReservationId)
	halpers.IfPanicError(err)

	roomReservation := r.RoomReservationService.FindById(ctx.Context(), int64(id))

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomReservation,
	}
	return ctx.JSON(webResponse)
}

func (r RoomReservationControllerImpl) Create(ctx *fiber.Ctx) error {
	var roomReservationCreateRequest request.RoomReservationCreateRequest
	err := ctx.BodyParser(&roomReservationCreateRequest)
	halpers.IfPanicError(err)

	file, err := ctx.FormFile("file")
	roomReservationCreateRequest.File = file

	roomReservation := r.RoomReservationService.Create(ctx.Context(), roomReservationCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomReservation,
	}
	return ctx.JSON(webResponse)
}

func (r RoomReservationControllerImpl) Update(ctx *fiber.Ctx) error {
	var roomReservationUpdateRequest request.RoomReservationUpdateRequest
	err := ctx.BodyParser(&roomReservationUpdateRequest)
	halpers.IfPanicError(err)

	roomReservationId := ctx.Params("roomReservationId")
	id, err := strconv.Atoi(roomReservationId)
	halpers.IfPanicError(err)

	roomReservationUpdateRequest.Id = int64(id)

	file, err := ctx.FormFile("file")
	roomReservationUpdateRequest.File = file

	roomReservation := r.RoomReservationService.Update(ctx.Context(), roomReservationUpdateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomReservation,
	}
	return ctx.JSON(webResponse)
}

func (r RoomReservationControllerImpl) Delete(ctx *fiber.Ctx) error {
	roomReservationId := ctx.Params("roomReservationId")
	id, err := strconv.Atoi(roomReservationId)
	halpers.IfPanicError(err)

	r.RoomReservationService.Delete(ctx.Context(), int64(id))

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}
	return ctx.JSON(webResponse)
}
