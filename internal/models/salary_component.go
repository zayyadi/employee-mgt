package models

import (
	"time"

	"github.com/google/uuid"
)

// SalaryComponent represents a component of an employee's salary (earning or deduction)
type SalaryComponent struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name" validate:"required"`
	Type        string    `gorm:"not null" json:"type" validate:"required,oneof=earning deduction"`
	IsTaxable   bool      `gorm:"default:true" json:"is_taxable"`
	IsRecurring bool      `gorm:"default:true" json:"is_recurring"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SalaryComponentCreate represents data for creating a new salary component
type SalaryComponentCreate struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required,oneof=earning deduction"`
	IsTaxable   bool   `json:"is_taxable"`
	IsRecurring bool   `json:"is_recurring"`
}

// SalaryComponentUpdate represents data for updating a salary component
type SalaryComponentUpdate struct {
	Name        string `json:"name"`
	Type        string `json:"type" validate:"oneof=earning deduction"`
	IsTaxable   bool   `json:"is_taxable"`
	IsRecurring bool   `json:"is_recurring"`
}

// SalaryComponentResponse represents salary component data returned in API responses
type SalaryComponentResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	IsTaxable   bool      `json:"is_taxable"`
	IsRecurring bool      `json:"is_recurring"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName specifies the table name for SalaryComponent model
func (SalaryComponent) TableName() string {
	return "salary_components"
}
