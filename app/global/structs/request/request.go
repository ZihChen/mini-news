package request

type RegisterUserOption struct {
	Username             string `json:"username" validate:"required,min=2,max=30"`
	Password             string `json:"password" validate:"required,min=6,max=20,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=6,max=20"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
