package models

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeSalary struct {
	ID                uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	EmployeeID        uuid.UUID  `gorm:"type:uuid;not null;index" json:"employee_id" validate:"required"`
	SalaryComponentID uuid.UUID  `gorm:"type:uuid;not null" json:"salary_component_id" validate:"required"`
	Amount            float64    `gorm:"not null" json:"amount" validate:"required,gt=0"`
	EffectiveDate     time.Time  `gorm:"not null" json:"effective_date" validate:"required"`
	EndDate           *time.Time `json:"end_date"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type EmployeeSalaryCreate struct {
	EmployeeID        uuid.UUID  `json:"employee_id" validate:"required"`
	SalaryComponentID uuid.UUID  `json:"salary_component_id" validate:"required"`
	Amount            float64    `json:"amount" validate:"required,gt=0"`
	EffectiveDate     time.Time  `json:"effective_date" validate:"required"`
	EndDate           *time.Time `json:"end_date"`
}

type EmployeeSalaryUpdate struct {
	Amount  float64    `json:"amount" validate:"gt=0"`
	EndDate *time.Time `json:"end_date"`
}

type EmployeeSalaryResponse struct {
	ID                uuid.UUID  `json:"id"`
	EmployeeID        uuid.UUID  `json:"employee_id"`
	SalaryComponentID uuid.UUID  `json:"salary_component_id"`
	Amount            float64    `json:"amount"`
	EffectiveDate     time.Time  `json:"effective_date"`
	EndDate           *time.Time `json:"end_date"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (EmployeeSalary) TableName() string {
	return "employee_salaries"
}
