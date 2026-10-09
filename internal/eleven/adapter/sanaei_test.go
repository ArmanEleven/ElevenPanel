package adapter

import "testing"

func TestReconcileReportsUnsupportedOperation(t *testing.T) {
	if err := NewSanaeiAdapter(nil, nil).Reconcile(); err == nil {
		t.Fatal("Reconcile returned nil before reconciliation is implemented")
	}
}
