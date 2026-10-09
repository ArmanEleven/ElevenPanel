package database

import (
	"path/filepath"
	"testing"

	elevenidentity "github.com/mhsanaei/3x-ui/v3/internal/eleven/identity"
	elevennodes "github.com/mhsanaei/3x-ui/v3/internal/eleven/nodes"
	elevenservice "github.com/mhsanaei/3x-ui/v3/internal/eleven/service"
)

func TestInitDBMigratesElevenTablesWithoutReplacingSanaeiTables(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "eleven-test.db")
	if err := InitDB(dbPath); err != nil {
		t.Fatalf("InitDB: %v", err)
	}
	t.Cleanup(func() {
		if err := CloseDB(); err != nil {
			t.Errorf("CloseDB: %v", err)
		}
	})

	db := GetDB()
	for _, table := range []string{
		"eleven_admins",
		"eleven_roles",
		"eleven_permissions",
		"eleven_users",
		"eleven_groups",
		"eleven_templates",
		"eleven_subscriptions",
		"eleven_nodes",
		"eleven_audit_events",
		"users",
		"inbounds",
	} {
		if !db.Migrator().HasTable(table) {
			t.Errorf("expected table %q to exist after migration", table)
		}
	}

	for _, index := range []struct {
		model any
		name  string
	}{
		{&elevenidentity.Admin{}, "ux_eleven_admin_username"},
		{&elevenidentity.Role{}, "ux_eleven_role_name"},
		{&elevenidentity.Permission{}, "ux_eleven_permission_key"},
		{&elevenservice.User{}, "ux_eleven_user_username"},
		{&elevennodes.Node{}, "ux_eleven_node_name"},
	} {
		if !db.Migrator().HasIndex(index.model, index.name) {
			t.Errorf("expected index %q on table %q", index.name, index.model)
		}
	}
}
