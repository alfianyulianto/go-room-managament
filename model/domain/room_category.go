package domain

import "time"

type RoomCategory struct {
	Id        int64
	Name      string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
