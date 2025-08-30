package request

type RegisterRequest struct {
	Name            string `validate:"required,max=255,uniqueUserNameCreate" json:"name"`
	Email           string `validate:"required,max=100,uniqueUserEmailCreate" json:"email"`
	Password        string `validate:"required,max=100" json:"password" form:"password"`
	ConfirmPassword string `validate:"required,max=100,eqfield=Password" json:"confirm_password" form:"confirm_password"`
	Phone           string `validate:"required,max=25" json:"phone"`
}
