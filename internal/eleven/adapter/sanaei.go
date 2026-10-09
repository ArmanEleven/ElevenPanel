package adapter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mhsanaei/3x-ui/v3/internal/database/model"
	"github.com/mhsanaei/3x-ui/v3/internal/web/service"
)

// SanaeiAdapter connects Eleven service-management operations
// to the existing Sanaei/3x-ui service layer.
//
// Eleven deliberately does not duplicate Xray/client creation logic.
// Sanaei remains responsible for protocol-specific defaults and
// applying the client to the requested inbound.
type SanaeiAdapter struct {
	clientService  *service.ClientService
	inboundService *service.InboundService
}

// NewSanaeiAdapter creates an adapter over the existing Sanaei services.
func NewSanaeiAdapter(
	clientService *service.ClientService,
	inboundService *service.InboundService,
) *SanaeiAdapter {
	if clientService == nil {
		clientService = &service.ClientService{}
	}

	if inboundService == nil {
		inboundService = &service.InboundService{}
	}

	return &SanaeiAdapter{
		clientService:  clientService,
		inboundService: inboundService,
	}
}

// splitAllowedIPs converts Sanaei's stored comma-separated string
// representation into the model.Client []string representation.
func splitAllowedIPs(value string) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}

	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if part != "" {
			result = append(result, part)
		}
	}

	return result
}

// ProvisionClient creates a client through Sanaei's existing service layer.
//
// Protocol-specific credentials are intentionally left to Sanaei.
func (a *SanaeiAdapter) ProvisionClient(input ProvisionClientInput) (ClientResult, error) {
	inboundID, err := strconv.Atoi(input.InboundID)
	if err != nil {
		return ClientResult{}, fmt.Errorf(
			"invalid inbound id %q: %w",
			input.InboundID,
			err,
		)
	}

	if strings.TrimSpace(input.Email) == "" {
		return ClientResult{}, fmt.Errorf("client email is required")
	}

	client := model.Client{
		ID:         input.UUID,
		Email:      input.Email,
		ExpiryTime: input.Expiry,
		TotalGB:    input.TotalGB,
		Enable:     true,
	}

	_, err = a.clientService.CreateOne(
		a.inboundService,
		inboundID,
		client,
	)
	if err != nil {
		return ClientResult{}, fmt.Errorf("provision client: %w", err)
	}

	record, err := a.clientService.GetRecordByEmail(nil, input.Email)
	if err != nil {
		return ClientResult{}, fmt.Errorf("read created client: %w", err)
	}

	return ClientResult{
		ID:        record.UUID,
		InboundID: input.InboundID,
		Email:     record.Email,
		Enabled:   record.Enable,
	}, nil
}

// UpdateClient accepts the returned UUID or a legacy email identifier.

func (a *SanaeiAdapter) findClientRecord(id string) (*model.ClientRecord, error) {
	if strings.TrimSpace(id) == "" {
		return nil, fmt.Errorf("client id is required")
	}
	if record, err := a.clientService.GetRecordByUUID(nil, id); err == nil {
		return record, nil
	}
	return a.clientService.GetRecordByEmail(nil, id)
}

func (a *SanaeiAdapter) UpdateClient(
	id string,
	input UpdateClientInput,
) (ClientResult, error) {
	if strings.TrimSpace(id) == "" {
		return ClientResult{}, fmt.Errorf("client id is required")
	}

	record, err := a.findClientRecord(id)
	if err != nil {
		return ClientResult{}, fmt.Errorf(
			"find client %q: %w",
			id,
			err,
		)
	}

	updated := model.Client{
		ID:              record.UUID,
		Email:           record.Email,
		Password:        record.Password,
		Auth:            record.Auth,
		Secret:          record.Secret,
		Flow:            record.Flow,
		Security:        record.Security,
		PrivateKey:      record.PrivateKey,
		PublicKey:       record.PublicKey,
		AllowedIPs:      splitAllowedIPs(record.AllowedIPs),
		PreSharedKey:    record.PreSharedKey,
		LimitIP:         record.LimitIP,
		TotalGB:         record.TotalGB,
		ExpiryTime:      record.ExpiryTime,
		Enable:          record.Enable,
		TgID:            record.TgID,
		Group:           record.Group,
		Comment:         record.Comment,
		Reset:           record.Reset,
		ResetDay:        record.ResetDay,
		ResetWeekday:    record.ResetWeekday,
		ResetMax:        record.ResetMax,
		TrafficReset:    record.TrafficReset,
		TrafficResetDay: record.TrafficResetDay,
		SubID:           record.SubID,
	}

	if input.Expiry != nil {
		updated.ExpiryTime = *input.Expiry
	}

	if input.TotalGB != nil {
		updated.TotalGB = *input.TotalGB
	}

	if input.Enabled != nil {
		updated.Enable = *input.Enabled
	}

	_, err = a.clientService.UpdateByEmail(
		a.inboundService,
		record.Email,
		updated,
		record.LimitHwid,
	)
	if err != nil {
		return ClientResult{}, fmt.Errorf(
			"update client %q: %w",
			id,
			err,
		)
	}

	result, err := a.clientService.GetRecordByEmail(nil, updated.Email)
	if err != nil {
		return ClientResult{}, fmt.Errorf(
			"read updated client: %w",
			err,
		)
	}

	return ClientResult{
		ID:      result.UUID,
		Email:   result.Email,
		Enabled: result.Enable,
	}, nil
}

// RevokeClient permanently removes the client through Sanaei's
// existing deletion flow.
//
// keepTraffic=false means traffic records are removed as part of
// the revoke operation.
func (a *SanaeiAdapter) RevokeClient(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("client id is required")
	}

	record, err := a.findClientRecord(id)
	if err != nil {
		return fmt.Errorf("find client %q: %w", id, err)
	}
	_, err = a.clientService.DeleteByEmail(
		a.inboundService,
		record.Email,
		false,
	)
	if err != nil {
		return fmt.Errorf(
			"revoke client %q: %w",
			id,
			err,
		)
	}

	return nil
}

func (a *SanaeiAdapter) GetClientTraffic(
	id string,
) (TrafficResult, error) {
	if strings.TrimSpace(id) == "" {
		return TrafficResult{}, fmt.Errorf("client id is required")
	}

	record, err := a.findClientRecord(id)
	if err != nil {
		return TrafficResult{}, fmt.Errorf("find client %q: %w", id, err)
	}
	traffic, err := a.inboundService.GetClientTrafficByEmail(record.Email)
	if err != nil {
		return TrafficResult{}, fmt.Errorf("get traffic for client %q: %w", id, err)
	}
	if traffic == nil {
		return TrafficResult{}, fmt.Errorf("traffic not found for client %q", id)
	}

	return TrafficResult{
		UpBytes:   traffic.Up,
		DownBytes: traffic.Down,
		Total:     traffic.Up + traffic.Down,
	}, nil
}

// GetInbound reads an existing Sanaei inbound without duplicating
// database logic.
func (a *SanaeiAdapter) GetInbound(id string) (InboundResult, error) {
	inboundID, err := strconv.Atoi(id)
	if err != nil {
		return InboundResult{}, fmt.Errorf(
			"invalid inbound id %q: %w",
			id,
			err,
		)
	}

	inbound, err := a.inboundService.GetInbound(inboundID)
	if err != nil {
		return InboundResult{}, fmt.Errorf(
			"get inbound %q: %w",
			id,
			err,
		)
	}

	return InboundResult{
		ID:       strconv.Itoa(inbound.Id),
		Protocol: string(inbound.Protocol),
		Remark:   inbound.Remark,
	}, nil
}

func (a *SanaeiAdapter) Reconcile() error {
	return fmt.Errorf("sanaei reconciliation not implemented")
}

// Compile-time assertion.
//
// If SanaeiAdapter stops implementing ClientAdapter, the project
// will fail at compile time instead of silently breaking the adapter.
var _ ClientAdapter = (*SanaeiAdapter)(nil)
