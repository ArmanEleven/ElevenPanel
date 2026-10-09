package adapter

import (
	"errors"
	"testing"

	"github.com/mhsanaei/3x-ui/v3/internal/config"
	"github.com/mhsanaei/3x-ui/v3/internal/database"
	"github.com/mhsanaei/3x-ui/v3/internal/database/dbtest"
	"github.com/mhsanaei/3x-ui/v3/internal/database/model"
	"github.com/mhsanaei/3x-ui/v3/internal/xray"
)

func TestReconcileReportsUnsupportedOperation(t *testing.T) {
	err := NewSanaeiAdapter(nil, nil).Reconcile()
	if !errors.Is(err, errReconciliationNotImplemented) {
		t.Fatalf("Reconcile() error = %v, want %v", err, errReconciliationNotImplemented)
	}
}

func TestClientIdentifierFallsBackToEmail(t *testing.T) {
	tests := []struct {
		name   string
		record model.ClientRecord
		want   string
	}{
		{name: "UUID", record: model.ClientRecord{UUID: "uuid-123", Email: "client@example.com"}, want: "uuid-123"},
		{name: "email fallback", record: model.ClientRecord{Email: "client@example.com"}, want: "client@example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clientIdentifier(&tt.record); got != tt.want {
				t.Fatalf("clientIdentifier() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGetClientTrafficUsesStableIDAndPreservesQuota(t *testing.T) {
	t.Setenv("XUI_DB_FOLDER", t.TempDir())
	dbtest.InitDB(t, config.GetDBPath())
	db := database.GetDB()

	record := model.ClientRecord{Email: "eleven-test@example.com", UUID: "uuid-123"}
	if err := db.Create(&record).Error; err != nil {
		t.Fatalf("create client record: %v", err)
	}
	traffic := xray.ClientTraffic{
		Email: record.Email,
		Up:    123,
		Down:  456,
		Total: 10000,
	}
	if err := db.Create(&traffic).Error; err != nil {
		t.Fatalf("create client traffic: %v", err)
	}

	got, err := NewSanaeiAdapter(nil, nil).GetClientTraffic("uuid-123")
	if err != nil {
		t.Fatalf("GetClientTraffic: %v", err)
	}
	if got.UpBytes != 123 || got.DownBytes != 456 {
		t.Fatalf("traffic = %+v, want upload 123 and download 456", got)
	}
	if got.LimitBytes != 10000 {
		t.Fatalf("quota = %d, want 10000", got.LimitBytes)
	}
}
