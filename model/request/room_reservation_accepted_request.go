package request

import (
	"mime/multipart"
)

type RoomReservationAcceptedRequest struct {
	Id      int64                 `validate:"required" json:"user_id" form:"user_id"`
	RoomId  int64                 `validate:"required,existRoom" json:"room_id" form:"room_id"`
	StartAt string                `validate:"required,datetime=2006-01-02 15:04:05,afterToday,overlap" json:"start_at" form:"start_at"`
	EndAt   string                `validate:"required,datetime=2006-01-02 15:04:05,after=StartAt" json:"end_at" form:"end_at"`
	Purpose string                `validate:"required,max=1000" json:"purpose" form:"purpose"`
	File    *multipart.FileHeader `validate:"required,file=pdf;docx;doc" json:"file" form:"file"`
}
