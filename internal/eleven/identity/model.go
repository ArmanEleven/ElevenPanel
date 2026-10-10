package identity

import "time"

// Admin is an Eleven operator account. Role and OwnerID are the basis for
// reseller scoping; authorization must still be enforced by the API layer.
type Admin struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username    string    `json:"username" gorm:"size:191;not null;uniqueIndex:ux_eleven_admin_username"`
	DisplayName string    `json:"displayName" gorm:"size:191;not null;default:''"`
	Role        string    `json:"role" gorm:"size:32;not null;index:idx_eleven_admin_owner_role"`
	OwnerID     *uint64   `json:"ownerId,omitempty" gorm:"index:idx_eleven_admin_owner_role"`
	Enabled     bool      `json:"enabled" gorm:"not null;default:true"`
	CreatedAt   time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"not null"`
}

func (Admin) TableName() string { return "eleven_admins" }

type Role struct {
	ID          uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"size:64;not null;uniqueIndex:ux_eleven_role_name"`
	Description string `json:"description" gorm:"size:255;not null;default:''"`
}

func (Role) TableName() string { return "eleven_roles" }

type Permission struct {
	ID          uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Key         string `json:"key" gorm:"size:128;not null;uniqueIndex:ux_eleven_permission_key"`
	Description string `json:"description" gorm:"size:255;not null;default:''"`
}

func (Permission) TableName() string { return "eleven_permissions" }

const (
	PermissionDashboardRead = "dashboard:read"
	PermissionAdminRead     = "admin:read"
	PermissionAdminWrite    = "admin:write"
	PermissionUserRead      = "user:read"
	PermissionUserWrite     = "user:write"
	PermissionNodeRead      = "node:read"
	PermissionNodeWrite     = "node:write"
	PermissionSubscription  = "subscription:write"
	PermissionSettingsWrite = "settings:write"
	PermissionAuditRead     = "audit:read"
)
