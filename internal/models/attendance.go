package models

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	EmployeeID   uuid.UUID `gorm:"type:uuid;not null;index" json:"employee_id" validate:"required"`
	CheckInTime  time.Time `gorm:"not null" json:"check_in_time" validate:"required"`
	CheckOutTime time.Time `json:"check_out_time"`
	Date         time.Time `gorm:"not null;index" json:"date" validate:"required"`
	Status       string    `gorm:"not null" json:"status" validate:"required,oneof=present late absent"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
}

type AttendanceCreate struct {
	EmployeeID  uuid.UUID `json:"employee_id" validate:"required"`
	CheckInTime time.Time `json:"check_in_time" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	Status      string    `json:"status" validate:"required,oneof=present late absent"`
	Notes       string    `json:"notes"`
}

type AttendanceUpdate struct {
	CheckOutTime *time.Time `json:"check_out_time"`
	Status       string     `json:"status" validate:"oneof=present late absent"`
	Notes        string     `json:"notes"`
}

type AttendanceResponse struct {
	ID           uuid.UUID `json:"id"`
	EmployeeID   uuid.UUID `json:"employee_id"`
	CheckInTime  time.Time `json:"check_in_time"`
	CheckOutTime time.Time `json:"check_out_time"`
	Date         time.Time `json:"date"`
	Status       string    `json:"status"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
}

func (Attendance) TableName() string {
	return "attendance"
}
