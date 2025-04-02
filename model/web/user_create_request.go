package web

type UserCreateRequest struct {
	Name            string `validate:"required" json:"name"`
	Username        string `validate:"required" json:"username"`
	Password        string `validate:"required" json:"password"`
	PasswordConfirm string `validate:"required" json:"password_confirm"`
}
