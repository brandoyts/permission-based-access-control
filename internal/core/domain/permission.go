package domain

import (
	"time"
)

type PermissionType string

const (
	PermissionCreate PermissionType = "can_create"
	PermissionRead   PermissionType = "can_read"
	PermissionUpdate PermissionType = "can_update"
	PermissionDelete PermissionType = "can_delete"
)

type Permission struct {
	Type      PermissionType
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToPermissionType(permission interface{}) PermissionType {
	val := permission.(string)
	return PermissionType(val)
}
