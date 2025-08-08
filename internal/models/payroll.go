package models

import (
	"time"

	"github.com/google/uuid"
)

// Payroll represents a payroll run for a specific period
type Payroll struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PayPeriodStart  time.Time `gorm:"not null;index" json:"pay_period_start" validate:"required"`
	PayPeriodEnd    time.Time `gorm:"not null;index" json:"pay_period_end" validate:"required"`
	PaymentDate     time.Time `gorm:"not null" json:"payment_date" validate:"required"`
	Status          string    `gorm:"not null;default:'draft'" json:"status" validate:"required,oneof=draft calculated approved processed"`
	TotalGrossPay   float64   `gorm:"not null;default:0" json:"total_gross_pay"`
	TotalDeductions float64   `gorm:"not null;default:0" json:"total_deductions"`
	TotalNetPay     float64   `gorm:"not null;default:0" json:"total_net_pay"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// PayrollCreate represents data for creating a new payroll run
type PayrollCreate struct {
	PayPeriodStart time.Time `json:"pay_period_start" validate:"required"`
	PayPeriodEnd   time.Time `json:"pay_period_end" validate:"required"`
	PaymentDate    time.Time `json:"payment_date" validate:"required"`
}

// PayrollUpdate represents data for updating a payroll run
type PayrollUpdate struct {
	Status          string  `json:"status" validate:"oneof=draft calculated approved processed"`
	TotalGrossPay   float64 `json:"total_gross_pay"`
	TotalDeductions float64 `json:"total_deductions"`
	TotalNetPay     float64 `json:"total_net_pay"`
}

// PayrollResponse represents payroll data returned in API responses
type PayrollResponse struct {
	ID              uuid.UUID `json:"id"`
	PayPeriodStart  time.Time `json:"pay_period_start"`
	PayPeriodEnd    time.Time `json:"pay_period_end"`
	PaymentDate     time.Time `json:"payment_date"`
	Status          string    `json:"status"`
	TotalGrossPay   float64   `json:"total_gross_pay"`
	TotalDeductions float64   `json:"total_deductions"`
	TotalNetPay     float64   `json:"total_net_pay"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// TableName specifies the table name for Payroll model
func (Payroll) TableName() string {
	return "payroll"
}
