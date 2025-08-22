package request

type UserUpdateRequest struct {
	Id    int64  `validate:"required" json:"id"`
	Name  string `validate:"required,max=255,uniqueUserNameUpdate" json:"name"`
	Email string `validate:"required,max=100,uniqueUserEmailUpdate" json:"email"`
	Phone string `validate:"required,max=25" json:"phone"`
	Level string `validate:"required,eq=Admin|eq=User" json:"level"`
}
