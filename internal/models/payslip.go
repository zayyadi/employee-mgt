package models

import (
	"time"

	"github.com/google/uuid"
)

// Payslip represents an employee's payslip for a specific payroll period
type Payslip struct {
	ID             uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	EmployeeID     uuid.UUID   `gorm:"type:uuid;not null;index" json:"employee_id" validate:"required"`
	PayrollID      uuid.UUID   `gorm:"type:uuid;not null;index" json:"payroll_id" validate:"required"`
	PayPeriodStart time.Time   `gorm:"not null" json:"pay_period_start" validate:"required"`
	PayPeriodEnd   time.Time   `gorm:"not null" json:"pay_period_end" validate:"required"`
	GrossPay       float64     `gorm:"not null" json:"gross_pay"`
	TaxAmount      float64     `gorm:"not null" json:"tax_amount"`
	Deductions     interface{} `gorm:"type:jsonb" json:"deductions"`
	NetPay         float64     `gorm:"not null" json:"net_pay"`
	FilePath       string      `json:"file_path"`
	CreatedAt      time.Time   `json:"created_at"`
}

// PayslipCreate represents data for creating a new payslip
type PayslipCreate struct {
	EmployeeID     uuid.UUID   `json:"employee_id" validate:"required"`
	PayrollID      uuid.UUID   `json:"payroll_id" validate:"required"`
	PayPeriodStart time.Time   `json:"pay_period_start" validate:"required"`
	PayPeriodEnd   time.Time   `json:"pay_period_end" validate:"required"`
	GrossPay       float64     `json:"gross_pay"`
	TaxAmount      float64     `json:"tax_amount"`
	Deductions     interface{} `json:"deductions"`
	NetPay         float64     `json:"net_pay"`
	FilePath       string      `json:"file_path"`
}

// PayslipUpdate represents data for updating a payslip
type PayslipUpdate struct {
	FilePath string `json:"file_path"`
}

// PayslipResponse represents payslip data returned in API responses
type PayslipResponse struct {
	ID             uuid.UUID   `json:"id"`
	EmployeeID     uuid.UUID   `json:"employee_id"`
	PayrollID      uuid.UUID   `json:"payroll_id"`
	PayPeriodStart time.Time   `json:"pay_period_start"`
	PayPeriodEnd   time.Time   `json:"pay_period_end"`
	GrossPay       float64     `json:"gross_pay"`
	TaxAmount      float64     `json:"tax_amount"`
	Deductions     interface{} `json:"deductions"`
	NetPay         float64     `json:"net_pay"`
	FilePath       string      `json:"file_path"`
	CreatedAt      time.Time   `json:"created_at"`
}

// TableName specifies the table name for Payslip model
func (Payslip) TableName() string {
	return "payslips"
}
