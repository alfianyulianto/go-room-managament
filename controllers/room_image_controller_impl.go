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

type RoomImageControllerImpl struct {
	services.RoomImageService
}

func NewRoomImageControllerImpl(roomImageService services.RoomImageService) *RoomImageControllerImpl {
	return &RoomImageControllerImpl{RoomImageService: roomImageService}
}

func (r RoomImageControllerImpl) FindAll(ctx *fiber.Ctx) error {
	roomIdStr := ctx.Params("roomId")
	roomId, err := strconv.Atoi(roomIdStr)
	halpers.IfPanicError(err)

	roomImages := r.RoomImageService.FindAll(ctx.Context(), int64(roomId))

	if roomImages == nil {
		roomImages = make([]web.RoomImageResponse, 0)
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomImages,
	}
	return ctx.JSON(webResponse)
}

func (r RoomImageControllerImpl) FindById(ctx *fiber.Ctx) error {
	roomIdStr := ctx.Params("roomId")
	roomId, err := strconv.Atoi(roomIdStr)
	halpers.IfPanicError(err)

	roomImageIdStr := ctx.Params("roomImageId")
	roomImageId, err := strconv.Atoi(roomImageIdStr)

	roomImage := r.RoomImageService.FindById(ctx.Context(), int64(roomId), int64(roomImageId))

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomImage,
	}
	return ctx.JSON(webResponse)
}

func (r RoomImageControllerImpl) Create(ctx *fiber.Ctx) error {
	var roomImageCreateRequest request.RoomImageCreateRequest

	roomIdStr := ctx.Params("roomId")
	roomId, err := strconv.Atoi(roomIdStr)
	halpers.IfPanicError(err)

	roomImageCreateRequest.RoomId = int64(roomId)

	image, err := ctx.FormFile("image")
	halpers.IfPanicError(err)
	roomImageCreateRequest.Image = image

	roomImage := r.RoomImageService.Create(ctx.Context(), roomImageCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomImage,
	}
	return ctx.JSON(webResponse)
}

func (r RoomImageControllerImpl) Update(ctx *fiber.Ctx) error {
	var roomImageUpdateRequest request.RoomImageUpdateRequest
	err := ctx.BodyParser(&roomImageUpdateRequest)
	halpers.IfPanicError(err)

	roomIdStr := ctx.Params("roomId")
	roomId, err := strconv.Atoi(roomIdStr)
	halpers.IfPanicError(err)

	roomImageUpdateRequest.RoomId = int64(roomId)
	roomImage := r.RoomImageService.Update(ctx.Context(), roomImageUpdateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomImage,
	}
	return ctx.JSON(webResponse)
}

func (r RoomImageControllerImpl) Delete(ctx *fiber.Ctx) error {
	roomIdStr := ctx.Params("roomId")
	roomId, err := strconv.Atoi(roomIdStr)
	halpers.IfPanicError(err)

	roomImageIdStr := ctx.Params("roomImageId")
	roomImageId, err := strconv.Atoi(roomImageIdStr)
	halpers.IfPanicError(err)

	r.RoomImageService.Delete(ctx.Context(), int64(roomId), int64(roomImageId))

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}
	return ctx.JSON(webResponse)
}
