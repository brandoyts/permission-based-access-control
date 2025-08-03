package domain

type User struct {
	ID          string           `json:"id,omitempty"`
	Permissions []PermissionType `json:"permissions,omitempty"`
	Email       string           `json:"email,omitempty"`
	Password    string           `json:"password,omitempty"`
	// CreatedAt   time.Time        `json:"created_at,omitempty"`
	// UpdatedAt   time.Time        `json:"updated_at,omitempty"`
	// DeletedAt   time.Time        `json:"deleted_at,omitempty"`
}
