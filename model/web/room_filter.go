package web

type RoomFilter struct {
	Search         *string `json:"name"`
	RoomCategoryId *string `json:"roomCategoryId"`
	Condition      *string `json:"condition"`
}
