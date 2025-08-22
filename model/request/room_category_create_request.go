package request

type RoomCategoryCreateRequest struct {
	Name string `validate:"required,max=255,uniqueRoomCategoryCreate" json:"name"`
}
