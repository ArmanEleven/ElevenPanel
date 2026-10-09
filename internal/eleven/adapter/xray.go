package adapter

// ClientAdapter isolates Eleven service management from Sanaei-specific client operations.
// SanaeiAdapter is the current implementation and keeps protocol logic in the existing core.
type ClientAdapter interface {
	ProvisionClient(input ProvisionClientInput) (ClientResult, error)
	UpdateClient(id string, input UpdateClientInput) (ClientResult, error)
	RevokeClient(id string) error
	GetClientTraffic(id string) (TrafficResult, error)
	GetInbound(id string) (InboundResult, error)
	Reconcile() error
}

type ProvisionClientInput struct {
	InboundID string
	Email     string
	UUID      string
	Expiry    int64
	TotalGB   int64
}

type UpdateClientInput struct {
	Expiry  *int64
	TotalGB *int64
	Enabled *bool
}

type ClientResult struct {
	ID        string
	InboundID string
	Email     string
	Enabled   bool
}

type TrafficResult struct {
	UpBytes   int64
	DownBytes int64
	LimitBytes int64
}

type InboundResult struct {
	ID       string
	Protocol string
	Remark   string
}
