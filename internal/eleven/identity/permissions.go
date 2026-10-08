package identity

var DefaultRoles = map[string][]string{
	"owner": {
		"dashboard:read",
		"admin:read",
		"admin:write",
		"user:read",
		"user:write",
		"node:read",
		"node:write",
		"subscription:write",
		"settings:write",
		"audit:read",
	},
	"manager": {
		"dashboard:read",
		"user:read",
		"user:write",
		"node:read",
		"node:write",
		"subscription:write",
		"audit:read",
	},
	"support": {
		"dashboard:read",
		"user:read",
	},
	"viewer": {
		"dashboard:read",
		"user:read",
		"node:read",
	},
}
