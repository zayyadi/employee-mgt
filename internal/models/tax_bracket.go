package models

import (
	"time"

	"github.com/google/uuid"
)

// TaxBracket represents a tax bracket for calculating employee taxes
type TaxBracket struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Country    string    `gorm:"not null;index" json:"country" validate:"required"`
	TaxYear    int       `gorm:"not null;index" json:"tax_year" validate:"required,min=2000,max=2100"`
	BracketMin float64   `gorm:"not null" json:"bracket_min" validate:"required,gte=0"`
	BracketMax float64   `gorm:"not null" json:"bracket_max" validate:"required,gte=0"`
	TaxRate    float64   `gorm:"not null" json:"tax_rate" validate:"required,gte=0,lte=100"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TaxBracketCreate represents data for creating a new tax bracket
type TaxBracketCreate struct {
	Country    string  `json:"country" validate:"required"`
	TaxYear    int     `json:"tax_year" validate:"required,min=2000,max=2100"`
	BracketMin float64 `json:"bracket_min" validate:"required,gte=0"`
	BracketMax float64 `json:"bracket_max" validate:"required,gte=0"`
	TaxRate    float64 `json:"tax_rate" validate:"required,gte=0,lte=100"`
}

// TaxBracketUpdate represents data for updating a tax bracket
type TaxBracketUpdate struct {
	Country    string  `json:"country"`
	TaxYear    int     `json:"tax_year" validate:"min=2000,max=2100"`
	BracketMin float64 `json:"bracket_min" validate:"gte=0"`
	BracketMax float64 `json:"bracket_max" validate:"gte=0"`
	TaxRate    float64 `json:"tax_rate" validate:"gte=0,lte=100"`
}

// TaxBracketResponse represents tax bracket data returned in API responses
type TaxBracketResponse struct {
	ID         uuid.UUID `json:"id"`
	Country    string    `json:"country"`
	TaxYear    int       `json:"tax_year"`
	BracketMin float64   `json:"bracket_min"`
	BracketMax float64   `json:"bracket_max"`
	TaxRate    float64   `json:"tax_rate"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName specifies the table name for TaxBracket model
func (TaxBracket) TableName() string {
	return "tax_brackets"
}
