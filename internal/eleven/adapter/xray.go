package adapter

// ClientAdapter is the narrow boundary between Eleven service management
// and the underlying Sanaei/3x-ui client implementation.
//
// The concrete implementation will be added after the pinned Sanaei source
// is vendored into the repository. Keeping this interface small prevents
// Eleven from duplicating Sanaei's Xray/client logic.
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
	Total     int64
}

type InboundResult struct {
	ID       string
	Protocol string
	Remark   string
}
