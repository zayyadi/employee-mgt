package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	EmployeeID  uuid.UUID `gorm:"type:uuid;not null;index" json:"employee_id" validate:"required"`
	Name        string    `gorm:"not null" json:"name" validate:"required"`
	Description string    `json:"description"`
	FilePath    string    `gorm:"not null" json:"file_path" validate:"required"`
	Category    string    `gorm:"not null" json:"category" validate:"required,oneof=contract certificate id_proof qualification other"`
	UploadedBy  uuid.UUID `gorm:"type:uuid;not null" json:"uploaded_by" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

type DocumentCreate struct {
	EmployeeID  uuid.UUID `json:"employee_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	FilePath    string    `json:"file_path" validate:"required"`
	Category    string    `json:"category" validate:"required,oneof=contract certificate id_proof qualification other"`
	UploadedBy  uuid.UUID `json:"uploaded_by" validate:"required"`
}

type DocumentUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category" validate:"oneof=contract certificate id_proof qualification other"`
}

type DocumentResponse struct {
	ID          uuid.UUID `json:"id"`
	EmployeeID  uuid.UUID `json:"employee_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FilePath    string    `json:"file_path"`
	Category    string    `json:"category"`
	UploadedBy  uuid.UUID `json:"uploaded_by"`
	CreatedAt   time.Time `json:"created_at"`
}

func (Document) TableName() string {
	return "documents"
}
