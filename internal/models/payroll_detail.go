package models

import (
	"time"

	"github.com/google/uuid"
)

// PayrollDetail represents the payroll details for a specific employee in a payroll run
type PayrollDetail struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PayrollID       uuid.UUID `gorm:"type:uuid;not null;index" json:"payroll_id" validate:"required"`
	EmployeeID      uuid.UUID `gorm:"type:uuid;not null;index" json:"employee_id" validate:"required"`
	GrossPay        float64   `gorm:"not null;default:0" json:"gross_pay"`
	TaxAmount       float64   `gorm:"not null;default:0" json:"tax_amount"`
	OtherDeductions float64   `gorm:"not null;default:0" json:"other_deductions"`
	NetPay          float64   `gorm:"not null;default:0" json:"net_pay"`
	CreatedAt       time.Time `json:"created_at"`
}

// PayrollDetailCreate represents data for creating a new payroll detail
type PayrollDetailCreate struct {
	PayrollID       uuid.UUID `json:"payroll_id" validate:"required"`
	EmployeeID      uuid.UUID `json:"employee_id" validate:"required"`
	GrossPay        float64   `json:"gross_pay"`
	TaxAmount       float64   `json:"tax_amount"`
	OtherDeductions float64   `json:"other_deductions"`
	NetPay          float64   `json:"net_pay"`
}

// PayrollDetailUpdate represents data for updating a payroll detail
type PayrollDetailUpdate struct {
	GrossPay        float64 `json:"gross_pay"`
	TaxAmount       float64 `json:"tax_amount"`
	OtherDeductions float64 `json:"other_deductions"`
	NetPay          float64 `json:"net_pay"`
}

// PayrollDetailResponse represents payroll detail data returned in API responses
type PayrollDetailResponse struct {
	ID              uuid.UUID `json:"id"`
	PayrollID       uuid.UUID `json:"payroll_id"`
	EmployeeID      uuid.UUID `json:"employee_id"`
	GrossPay        float64   `json:"gross_pay"`
	TaxAmount       float64   `json:"tax_amount"`
	OtherDeductions float64   `json:"other_deductions"`
	NetPay          float64   `json:"net_pay"`
	CreatedAt       time.Time `json:"created_at"`
}

// TableName specifies the table name for PayrollDetail model
func (PayrollDetail) TableName() string {
	return "payroll_details"
}
