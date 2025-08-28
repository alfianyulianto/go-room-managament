package request

type RoomUpdateRequest struct {
	Id             int64  `validate:"required" json:"id"`
	RoomCategoryId int64  `validate:"required,existRoomCategory" json:"room_category_id" form:"room_category_id"`
	Code           string `validate:"required,max=100,uniqueRoomCodeUpdate" json:"code"`
	Name           string `validate:"required,max=255,uniqueRoomNameUpdate" json:"name"`
	Condition      string `validate:"required,eq=Baik|eq=Rusak Ringan|eq=Rusak Sedang|eq=Rusak Berat" json:"condition"`
	Note           string `validate:"required,max=1000" json:"note"`
}
