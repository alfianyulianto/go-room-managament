package web

import "time"

type RoomImageResponse struct {
	Id        int64     `json:"id"`
	RoomId    int64     `json:"room_id"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
