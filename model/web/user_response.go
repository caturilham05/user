package web

import "time"

type UserResponse struct {
	Id        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Username  string    `json:"username" example:"johndoe"`
	CreatedAt time.Time `json:"created_at" example:"2025-04-03T01:13:12Z"`
}
