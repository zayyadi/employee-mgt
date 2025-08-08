package models

import (
	"time"

	"github.com/google/uuid"
)

// Report represents a report that can be generated in the system
type Report struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string    `gorm:"not null" json:"name" validate:"required"`
	Type        string    `gorm:"not null;index" json:"type" validate:"required"`
	Description string    `json:"description"`
	Query       string    `gorm:"not null" json:"query" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

// ReportCreate represents data for creating a new report
type ReportCreate struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Description string `json:"description"`
	Query       string `json:"query" validate:"required"`
}

// ReportUpdate represents data for updating a report
type ReportUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Query       string `json:"query"`
}

// ReportResponse represents report data returned in API responses
type ReportResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Query       string    `json:"query"`
	CreatedAt   time.Time `json:"created_at"`
}

// TableName specifies the table name for Report model
func (Report) TableName() string {
	return "reports"
}
