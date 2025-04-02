package web

type UserCreateRequest struct {
	Name            string `validate:"required" json:"name" example:"John Doe"`
	Username        string `validate:"required" json:"username" example:"johndoe"`
	Password        string `validate:"required" json:"password" example:"123456"`
	PasswordConfirm string `validate:"required" json:"password_confirm" example:"123456"`
}
