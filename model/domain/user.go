package domain

import "time"

type User struct {
	Id              int
	Name            string
	Username        string
	Password        string
	PasswordConfirm string
	CreatedAt       time.Time
	Token           string
	// CreatedAt sql.NullTime
	// CreatedAt string
}
