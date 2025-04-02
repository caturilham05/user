package web

type UserUpdateRequest struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Username        string `validate:"required" json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
