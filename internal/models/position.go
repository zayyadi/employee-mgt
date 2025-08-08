package models

import (
	"time"

	"github.com/google/uuid"
)

// Position represents a job position in the organization
type Position struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Title          string    `gorm:"not null" json:"title" validate:"required"`
	DepartmentID   uuid.UUID `gorm:"type:uuid;not null" json:"department_id" validate:"required"`
	Description    string    `gorm:"not null" json:"description" validate:"required"`
	Requirements   string    `gorm:"not null" json:"requirements" validate:"required"`
	SalaryRangeMin float64   `gorm:"not null" json:"salary_range_min" validate:"required,gt=0"`
	SalaryRangeMax float64   `gorm:"not null" json:"salary_range_max" validate:"required,gt=0"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// PositionCreate represents data for creating a new position
type PositionCreate struct {
	Title          string    `json:"title" validate:"required"`
	DepartmentID   uuid.UUID `json:"department_id" validate:"required"`
	Description    string    `json:"description" validate:"required"`
	Requirements   string    `json:"requirements" validate:"required"`
	SalaryRangeMin float64   `json:"salary_range_min" validate:"required,gt=0"`
	SalaryRangeMax float64   `json:"salary_range_max" validate:"required,gt=0"`
}

// PositionUpdate represents data for updating a position
type PositionUpdate struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Requirements   string  `json:"requirements"`
	SalaryRangeMin float64 `json:"salary_range_min" validate:"gt=0"`
	SalaryRangeMax float64 `json:"salary_range_max" validate:"gt=0"`
}

// PositionResponse represents position data returned in API responses
type PositionResponse struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	DepartmentID   uuid.UUID `json:"department_id"`
	Description    string    `json:"description"`
	Requirements   string    `json:"requirements"`
	SalaryRangeMin float64   `json:"salary_range_min"`
	SalaryRangeMax float64   `json:"salary_range_max"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// TableName specifies the table name for Position model
func (Position) TableName() string {
	return "positions"
}
