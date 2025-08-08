package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID                    uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID                uuid.UUID  `gorm:"type:uuid;not null" json:"user_id" validate:"required"`
	EmployeeID            string     `gorm:"uniqueIndex;not null" json:"employee_id" validate:"required"`
	FirstName             string     `gorm:"not null" json:"first_name" validate:"required"`
	LastName              string     `gorm:"not null" json:"last_name" validate:"required"`
	DateOfBirth           time.Time  `gorm:"not null" json:"date_of_birth" validate:"required"`
	Gender                string     `gorm:"not null" json:"gender" validate:"required,oneof=male female other"`
	MaritalStatus         string     `gorm:"not null" json:"marital_status" validate:"required,oneof=single married divorced widowed"`
	PhoneNumber           string     `gorm:"not null" json:"phone_number" validate:"required"`
	Email                 string     `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	Address               string     `gorm:"not null" json:"address" validate:"required"`
	EmergencyContactName  string     `gorm:"not null" json:"emergency_contact_name" validate:"required"`
	EmergencyContactPhone string     `gorm:"not null" json:"emergency_contact_phone" validate:"required"`
	DepartmentID          uuid.UUID  `gorm:"type:uuid" json:"department_id"`
	PositionID            uuid.UUID  `gorm:"type:uuid" json:"position_id"`
	HireDate              time.Time  `gorm:"not null" json:"hire_date" validate:"required"`
	EmploymentStatus      string     `gorm:"not null" json:"employment_status" validate:"required,oneof=active inactive terminated"`
	ManagerID             *uuid.UUID `gorm:"type:uuid" json:"manager_id"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

type EmployeeCreate struct {
	UserID                uuid.UUID  `json:"user_id" validate:"required"`
	EmployeeID            string     `json:"employee_id" validate:"required"`
	FirstName             string     `json:"first_name" validate:"required"`
	LastName              string     `json:"last_name" validate:"required"`
	DateOfBirth           time.Time  `json:"date_of_birth" validate:"required"`
	Gender                string     `json:"gender" validate:"required,oneof=male female other"`
	MaritalStatus         string     `json:"marital_status" validate:"required,oneof=single married divorced widowed"`
	PhoneNumber           string     `json:"phone_number" validate:"required"`
	Email                 string     `json:"email" validate:"required,email"`
	Address               string     `json:"address" validate:"required"`
	EmergencyContactName  string     `json:"emergency_contact_name" validate:"required"`
	EmergencyContactPhone string     `json:"emergency_contact_phone" validate:"required"`
	DepartmentID          *uuid.UUID `json:"department_id"`
	PositionID            *uuid.UUID `json:"position_id"`
	HireDate              time.Time  `json:"hire_date" validate:"required"`
	EmploymentStatus      string     `json:"employment_status" validate:"required,oneof=active inactive terminated"`
	ManagerID             *uuid.UUID `json:"manager_id"`
}

type EmployeeUpdate struct {
	FirstName             string     `json:"first_name"`
	LastName              string     `json:"last_name"`
	DateOfBirth           *time.Time `json:"date_of_birth"`
	Gender                string     `json:"gender" validate:"oneof=male female other"`
	MaritalStatus         string     `json:"marital_status" validate:"oneof=single married divorced widowed"`
	PhoneNumber           string     `json:"phone_number"`
	Email                 string     `json:"email" validate:"email"`
	Address               string     `json:"address"`
	EmergencyContactName  string     `json:"emergency_contact_name"`
	EmergencyContactPhone string     `json:"emergency_contact_phone"`
	DepartmentID          *uuid.UUID `json:"department_id"`
	PositionID            *uuid.UUID `json:"position_id"`
	HireDate              *time.Time `json:"hire_date"`
	EmploymentStatus      string     `json:"employment_status" validate:"oneof=active inactive terminated"`
	ManagerID             *uuid.UUID `json:"manager_id"`
}

type EmployeeResponse struct {
	ID                    uuid.UUID  `json:"id"`
	UserID                uuid.UUID  `json:"user_id"`
	EmployeeID            string     `json:"employee_id"`
	FirstName             string     `json:"first_name"`
	LastName              string     `json:"last_name"`
	DateOfBirth           time.Time  `json:"date_of_birth"`
	Gender                string     `json:"gender"`
	MaritalStatus         string     `json:"marital_status"`
	PhoneNumber           string     `json:"phone_number"`
	Email                 string     `json:"email"`
	Address               string     `json:"address"`
	EmergencyContactName  string     `json:"emergency_contact_name"`
	EmergencyContactPhone string     `json:"emergency_contact_phone"`
	DepartmentID          *uuid.UUID `json:"department_id"`
	PositionID            *uuid.UUID `json:"position_id"`
	HireDate              time.Time  `json:"hire_date"`
	EmploymentStatus      string     `json:"employment_status"`
	ManagerID             *uuid.UUID `json:"manager_id"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

func (Employee) TableName() string {
	return "employees"
}
