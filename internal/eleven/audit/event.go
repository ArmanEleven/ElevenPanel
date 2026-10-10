package audit

import "time"

// Event is an append-only audit record. Application code should create events
// through the audit service and must not expose update/delete operations.
type Event struct {
	ID         uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	ActorID    *uint64   `json:"actorId,omitempty" gorm:"index:idx_eleven_audit_actor_time"`
	Action     string    `json:"action" gorm:"size:128;not null;index"`
	Resource   string    `json:"resource" gorm:"size:128;not null;index"`
	ResourceID string    `json:"resourceId,omitempty" gorm:"size:191;not null;default:''"`
	Metadata   string    `json:"metadata,omitempty" gorm:"type:text;not null;default:''"`
	CreatedAt  time.Time `json:"createdAt" gorm:"not null;index:idx_eleven_audit_actor_time"`
}

func (Event) TableName() string { return "eleven_audit_events" }
