package web

import "time"

type UserResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	// CreatedAt sql.NullTime `json:"created_at"`
	// CreatedAt string `json:"created_at"`
}
