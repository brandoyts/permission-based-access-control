package domain

import "time"

type User struct {
	ID          string
	Permissions []Permission
	Email       string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
