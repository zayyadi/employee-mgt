package models

import (
	"time"

	"github.com/google/uuid"
)

type Department struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string     `gorm:"uniqueIndex;not null" json:"name" validate:"required"`
	Description string     `gorm:"not null" json:"description" validate:"required"`
	ManagerID   *uuid.UUID `gorm:"type:uuid" json:"manager_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type DepartmentCreate struct {
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description" validate:"required"`
	ManagerID   *uuid.UUID `json:"manager_id"`
}

type DepartmentUpdate struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ManagerID   *uuid.UUID `json:"manager_id"`
}

type DepartmentResponse struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ManagerID   *uuid.UUID `json:"manager_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (Department) TableName() string {
	return "departments"
}
