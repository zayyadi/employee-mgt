package models

import (
	"time"

	"github.com/google/uuid"
)

type LeaveRequest struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	EmployeeID  uuid.UUID  `gorm:"type:uuid;not null;index" json:"employee_id" validate:"required"`
	LeaveTypeID uuid.UUID  `gorm:"type:uuid;not null" json:"leave_type_id" validate:"required"`
	StartDate   time.Time  `gorm:"not null" json:"start_date" validate:"required"`
	EndDate     time.Time  `gorm:"not null" json:"end_date" validate:"required"`
	Reason      string     `gorm:"not null" json:"reason" validate:"required"`
	Status      string     `gorm:"not null;default:'pending'" json:"status" validate:"required,oneof=pending approved rejected"`
	ApprovedBy  *uuid.UUID `gorm:"type:uuid" json:"approved_by"`
	ApprovedAt  *time.Time `json:"approved_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type LeaveRequestCreate struct {
	EmployeeID  uuid.UUID `json:"employee_id" validate:"required"`
	LeaveTypeID uuid.UUID `json:"leave_type_id" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Reason      string    `json:"reason" validate:"required"`
}

type LeaveRequestUpdate struct {
	Status     string     `json:"status" validate:"oneof=pending approved rejected"`
	ApprovedBy *uuid.UUID `json:"approved_by"`
	ApprovedAt *time.Time `json:"approved_at"`
}

type LeaveRequestResponse struct {
	ID          uuid.UUID  `json:"id"`
	EmployeeID  uuid.UUID  `json:"employee_id"`
	LeaveTypeID uuid.UUID  `json:"leave_type_id"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     time.Time  `json:"end_date"`
	Reason      string     `json:"reason"`
	Status      string     `json:"status"`
	ApprovedBy  *uuid.UUID `json:"approved_by"`
	ApprovedAt  *time.Time `json:"approved_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (LeaveRequest) TableName() string {
	return "leave_requests"
}
