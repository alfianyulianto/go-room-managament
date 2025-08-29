package web

import (
	"time"
)

type RoomReservationResponse struct {
	Id        int64         `json:"id"`
	UserId    int64         `json:"user_id"`
	UserName  *string       `json:"user_name,omitempty"`
	RoomId    int64         `json:"room_id"`
	RoomName  string        `json:"room_name,omitempty"`
	StarAt    string        `json:"star_at"`
	EndAt     string        `json:"end_at"`
	Purpose   string        `json:"purpose"`
	Status    string        `json:"status"`
	ApproveId *int64        `json:"approve_id"`
	File      string        `json:"file"`
	User      *UserResponse `json:"user"`
	Room      *RoomResponse `json:"room"`
	Approve   *UserResponse `json:"approve"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
