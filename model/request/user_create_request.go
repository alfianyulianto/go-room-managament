package request

type UserCreateRequest struct {
	Name  string `validate:"required,max=255,uniqueUserNameCreate" json:"name"`
	Email string `validate:"required,max=100,uniqueUserEmailCreate" json:"email"`
	Phone string `validate:"required,max=25" json:"phone"`
	Level string `validate:"required,eq=Admin|eq=User" json:"level"`
}
