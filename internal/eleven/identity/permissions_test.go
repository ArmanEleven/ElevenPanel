package identity

import (
	"reflect"
	"testing"
)

func TestPermissionsForRoleReturnsCopy(t *testing.T) {
	got := PermissionsForRole("owner")
	if len(got) == 0 {
		t.Fatal("owner should have permissions")
	}
	got[0] = "privilege:invented"
	if HasPermission("owner", "privilege:invented") {
		t.Fatal("mutating returned permissions must not change the policy")
	}
}

func TestRolePermissions(t *testing.T) {
	tests := []struct {
		role       string
		permission string
		want       bool
	}{
		{"owner", "settings:write", true},
		{"reseller", "user:write", true},
		{"reseller", "settings:write", false},
		{"reseller", "admin:write", false},
		{"customer", "subscription:read", true},
		{"customer", "subscription:write", false},
		{"customer", "node:read", false},
		{"unknown", "dashboard:read", false},
	}
	for _, tt := range tests {
		t.Run(tt.role+"/"+tt.permission, func(t *testing.T) {
			if got := HasPermission(tt.role, tt.permission); got != tt.want {
				t.Fatalf("HasPermission(%q, %q) = %v, want %v", tt.role, tt.permission, got, tt.want)
			}
		})
	}
}

func TestPermissionsForRole(t *testing.T) {
	got := PermissionsForRole("customer")
	want := []string{"dashboard:read", "subscription:read", "profile:read"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("PermissionsForRole(customer) = %#v, want %#v", got, want)
	}
	if got := PermissionsForRole("missing-role"); got != nil {
		t.Fatalf("unknown role permissions = %#v, want nil", got)
	}
}
