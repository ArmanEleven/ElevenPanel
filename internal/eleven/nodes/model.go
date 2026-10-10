package nodes

import "time"

type Node struct {
	ID       uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string    `json:"name" gorm:"size:191;not null;uniqueIndex:ux_eleven_node_name"`
	Address  string    `json:"address" gorm:"size:255;not null"`
	Port     uint16    `json:"port" gorm:"not null;default:2053"`
	Enabled  bool      `json:"enabled" gorm:"not null;default:true"`
	LastSeen time.Time `json:"lastSeen" gorm:"not null"`
}

func (Node) TableName() string { return "eleven_nodes" }

type HealthStatus string

const (
	HealthUnknown HealthStatus = "unknown"
	HealthOnline  HealthStatus = "online"
	HealthOffline HealthStatus = "offline"
)
