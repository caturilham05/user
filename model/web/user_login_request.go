package web

type UserLoginRequest struct {
	Username string `validate:"required" json:"username" example:"johndoe"`
	Password string `validate:"required" json:"password" example:"password123"`
}
