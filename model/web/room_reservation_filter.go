package web

type RoomReservationFilter struct {
	Search  *string `json:"name"`
	StartAt *string `json:"startAt"`
	EndAt   *string `json:"endAt"`
	Status  *string `json:"status"`
}
