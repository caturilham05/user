package web

type UserUpdateRequest struct {
	Id              int    `json:"id" example:"1"`
	Name            string `json:"name" example:"John Doe"`
	Username        string `validate:"required" example:"johndoe"`
	Password        string `json:"password" example:"123456"`
	PasswordConfirm string `json:"password_confirm" example:"123456"`
}
