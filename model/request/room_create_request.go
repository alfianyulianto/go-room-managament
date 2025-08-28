package request

import "mime/multipart"

type RoomCreateRequest struct {
	RoomCategoryId int64                   `validate:"required,existRoomCategory" json:"room_category_id" form:"room_category_id"`
	Code           string                  `validate:"required,max=100,uniqueRoomCodeCreate" json:"code" form:"code"`
	Name           string                  `validate:"required,max=255,uniqueRoomNameCreate" json:"name" form:"name"`
	Condition      string                  `validate:"required,eq=Baik|eq=Rusak Ringan|eq=Rusak Sedang|eq=Rusak Berat" json:"condition"`
	Note           string                  `validate:"required,max=1000" json:"note"`
	Images         []*multipart.FileHeader `validate:"required,files" json:"images"`
}
