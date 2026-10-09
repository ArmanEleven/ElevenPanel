package service

import "time"

type User struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username    string    `json:"username" gorm:"size:191;not null;uniqueIndex:ux_eleven_user_username"`
	DisplayName string    `json:"displayName" gorm:"size:191;not null;default:''"`
	Enabled     bool      `json:"enabled" gorm:"not null;default:true"`
	OwnerID     *uint64   `json:"ownerId,omitempty" gorm:"index:idx_eleven_user_owner"`
	GroupID     *uint64   `json:"groupId,omitempty" gorm:"index:idx_eleven_user_group"`
	CreatedAt   time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"not null"`
}

func (User) TableName() string { return "eleven_users" }

type Group struct {
	ID          uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"size:191;not null;uniqueIndex:ux_eleven_group_name"`
	Description string `json:"description" gorm:"size:255;not null;default:''"`
}

func (Group) TableName() string { return "eleven_groups" }

type Template struct {
	ID           uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name" gorm:"size:191;not null;uniqueIndex:ux_eleven_template_name"`
	Description  string `json:"description" gorm:"size:255;not null;default:''"`
	TotalGB      int64  `json:"totalGB" gorm:"not null;default:0"`
	DurationDays int    `json:"durationDays" gorm:"not null;default:30"`
}

func (Template) TableName() string { return "eleven_templates" }

type Subscription struct {
	ID         uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     uint64    `json:"userId" gorm:"not null;index:idx_eleven_subscription_user"`
	TemplateID *uint64   `json:"templateId,omitempty" gorm:"index"`
	ExpiresAt  time.Time `json:"expiresAt" gorm:"not null;index"`
	Enabled    bool      `json:"enabled" gorm:"not null;default:true"`
	TotalGB    int64     `json:"totalGB" gorm:"not null;default:0"`
	CreatedAt  time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"not null"`
}

func (Subscription) TableName() string { return "eleven_subscriptions" }
