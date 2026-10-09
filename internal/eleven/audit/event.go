package audit

import "time"

type Event struct {
	ID         uint64    `json:"id"`
	ActorID    *uint64   `json:"actorId,omitempty"`
	Action     string    `json:"action"`
	Resource   string    `json:"resource"`
	ResourceID string    `json:"resourceId,omitempty"`
	Metadata   string    `json:"metadata,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}
