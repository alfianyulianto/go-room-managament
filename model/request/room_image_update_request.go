package request

import "mime/multipart"

type RoomImageUpdateRequest struct {
	Id     int64                 `json:"id" from:"id"`
	RoomId int64                 `validate:"required" json:"room_id" from:"room_id"`
	Image  *multipart.FileHeader `validate:"required,file" json:"image" form:"image"`
}
