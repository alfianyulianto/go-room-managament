package request

import "mime/multipart"

type RoomImageCreateRequest struct {
	RoomId int64                 `validate:"required" json:"room_id" form:"room_id"`
	Image  *multipart.FileHeader `validate:"required,file" json:"image" form:"image"`
}
