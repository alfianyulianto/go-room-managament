package web

type UserFilter struct {
	Search *string `json:"name"`
	Level  *string `json:"level"`
}
