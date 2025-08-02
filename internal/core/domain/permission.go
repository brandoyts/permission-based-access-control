package domain

import "time"

type Permission struct {
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
