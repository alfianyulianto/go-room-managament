package domain

import (
	"database/sql"
	"time"
)

type RoomReservation struct {
	Id          int64          `json:"id"`
	UserId      int64          `json:"user_id"`
	UserName    string         `json:"user_name"`
	RoomId      int64          `json:"room_id"`
	RoomName    string         `json:"room_name"`
	StartAt     time.Time      `json:"start_at"`
	EndAt       time.Time      `json:"end_at"`
	Purpose     string         `json:"purpose"`
	Status      string         `json:"status"`
	ApproveId   sql.NullInt64  `json:"approve_id"`
	ApproveName sql.NullString `json:"approve_name"`
	File        string         `json:"file"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
