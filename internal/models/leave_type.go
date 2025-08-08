package models

import (
	"time"

	"github.com/google/uuid"
)

type LeaveType struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name           string    `gorm:"uniqueIndex;not null" json:"name" validate:"required"`
	Description    string    `gorm:"not null" json:"description" validate:"required"`
	MaxDaysPerYear int       `gorm:"not null" json:"max_days_per_year" validate:"required,min=0"`
	IsAccrued      bool      `gorm:"default:false" json:"is_accrued"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type LeaveTypeCreate struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description" validate:"required"`
	MaxDaysPerYear int    `json:"max_days_per_year" validate:"required,min=0"`
	IsAccrued      bool   `json:"is_accrued"`
}

type LeaveTypeUpdate struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	MaxDaysPerYear int    `json:"max_days_per_year" validate:"min=0"`
	IsAccrued      bool   `json:"is_accrued"`
}

type LeaveTypeResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	MaxDaysPerYear int       `json:"max_days_per_year"`
	IsAccrued      bool      `json:"is_accrued"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (LeaveType) TableName() string {
	return "leave_types"
}
