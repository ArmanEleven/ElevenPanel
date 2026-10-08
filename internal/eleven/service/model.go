package service

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	DisplayName string  `json:"displayName"`
	Enabled   bool      `json:"enabled"`
	GroupID   *uint64   `json:"groupId,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Group struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

type Template struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

type Subscription struct {
	ID         uint64    `json:"id"`
	UserID     uint64    `json:"userId"`
	TemplateID *uint64   `json:"templateId,omitempty"`
	ExpiresAt  time.Time `json:"expiresAt"`
	Enabled    bool      `json:"enabled"`
}
