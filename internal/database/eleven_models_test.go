package database

import (
	"path/filepath"
	"testing"
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

	for _, index := range []struct { table, name string }{
		{"eleven_admins", "ux_eleven_admin_username"},
		{"eleven_roles", "ux_eleven_role_name"},
		{"eleven_permissions", "ux_eleven_permission_key"},
		{"eleven_users", "ux_eleven_user_username"},
		{"eleven_nodes", "ux_eleven_node_name"},
	} {
		if !db.Migrator().HasIndex(index.table, index.name) {
			t.Errorf("expected index %q on table %q", index.name, index.table)
		}
	}
}
