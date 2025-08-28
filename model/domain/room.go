package domain

import "time"

type Room struct {
	Id               int64     `json:"id"`
	RoomCategoryId   int64     `json:"room_category_id"`
	RoomCategoryName string    `json:"room_category_name"`
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	Condition        string    `json:"condition"`
	Note             string    `json:"note"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
