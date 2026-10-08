package identity

import "time"

type Admin struct {
	ID           uint64    `json:"id"`
	Username     string    `json:"username"`
	DisplayName  string    `json:"displayName"`
	Enabled      bool      `json:"enabled"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Role struct {
	ID          uint64   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
}

type Permission struct {
	ID          uint64   `json:"id"`
	Key         string   `json:"key"`
	Description string   `json:"description"`
}

const (
	PermissionDashboardRead = "dashboard:read"
	PermissionAdminRead     = "admin:read"
	PermissionAdminWrite    = "admin:write"
	PermissionUserRead      = "user:read"
	PermissionUserWrite     = "user:write"
	PermissionNodeRead      = "node:read"
	PermissionNodeWrite     = "node:write"
	PermissionSubscription = "subscription:write"
	PermissionSettingsWrite = "settings:write"
	PermissionAuditRead     = "audit:read"
)
