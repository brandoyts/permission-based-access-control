package domain

import "time"

type PermissionType string

type Permission struct {
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
