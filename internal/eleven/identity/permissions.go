package identity

// DefaultRoles is the server-side permission policy for Eleven identities.
// Callers should use PermissionsForRole so the shared policy cannot be mutated
// through a returned slice.
var DefaultRoles = map[string][]string{
	"owner": {
		"dashboard:read",
		"admin:read",
		"admin:write",
		"user:read",
		"user:write",
		"node:read",
		"node:write",
		"subscription:read",
		"subscription:write",
		"settings:write",
		"audit:read",
		"profile:read",
	},
	// Resellers can manage customers and subscriptions, but cannot manage
	// operators, server nodes, or global settings.
	"reseller": {
		"dashboard:read",
		"user:read",
		"user:write",
		"subscription:read",
		"subscription:write",
		"profile:read",
	},
	// Customers are restricted to their own account and subscription views.
	// Resource-level ownership checks are still required on future data routes.
	"customer": {
		"dashboard:read",
		"subscription:read",
		"profile:read",
	},
	// Legacy roles remain recognized for compatibility with early Eleven data.
	"manager": {
		"dashboard:read",
		"user:read",
		"user:write",
		"node:read",
		"node:write",
		"subscription:read",
		"subscription:write",
		"audit:read",
		"profile:read",
	},
	"support": {
		"dashboard:read",
		"user:read",
		"profile:read",
	},
	"viewer": {
		"dashboard:read",
		"user:read",
		"node:read",
		"profile:read",
	},
}

// PermissionsForRole returns a defensive copy of a role's permission set.
func PermissionsForRole(role string) []string {
	permissions, ok := DefaultRoles[role]
	if !ok {
		return nil
	}
	return append([]string(nil), permissions...)
}

// HasPermission reports whether a known role grants the requested permission.
// Unknown roles fail closed.
func HasPermission(role, permission string) bool {
	for _, candidate := range DefaultRoles[role] {
		if candidate == permission {
			return true
		}
	}
	return false
}
