package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id" validate:"required"`
	Title     string    `gorm:"not null" json:"title" validate:"required"`
	Message   string    `gorm:"not null" json:"message" validate:"required"`
	IsRead    bool      `gorm:"default:false;index" json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationCreate struct {
	UserID  uuid.UUID `json:"user_id" validate:"required"`
	Title   string    `json:"title" validate:"required"`
	Message string    `json:"message" validate:"required"`
}

type NotificationUpdate struct {
	IsRead bool `json:"is_read"`
}

type NotificationResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

func (Notification) TableName() string {
	return "notifications"
}
