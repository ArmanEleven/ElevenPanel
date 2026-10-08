package nodes

import "time"

type Node struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Port      uint16    `json:"port"`
	Enabled   bool      `json:"enabled"`
	LastSeen  time.Time `json:"lastSeen"`
}

type HealthStatus string

const (
	HealthUnknown HealthStatus = "unknown"
	HealthOnline  HealthStatus = "online"
	HealthOffline HealthStatus = "offline"
)
