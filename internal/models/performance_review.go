package models

import (
	"time"

	"github.com/google/uuid"
)

// PerformanceReview represents a performance review for an employee
type PerformanceReview struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	EmployeeID uuid.UUID `gorm:"type:uuid;not null;index" json:"employee_id" validate:"required"`
	ReviewerID uuid.UUID `gorm:"type:uuid;not null" json:"reviewer_id" validate:"required"`
	ReviewDate time.Time `gorm:"not null" json:"review_date" validate:"required"`
	Rating     int       `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating" validate:"required,min=1,max=5"`
	Comments   string    `gorm:"not null" json:"comments" validate:"required"`
	Goals      string    `json:"goals"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// PerformanceReviewCreate represents data for creating a new performance review
type PerformanceReviewCreate struct {
	EmployeeID uuid.UUID `json:"employee_id" validate:"required"`
	ReviewerID uuid.UUID `json:"reviewer_id" validate:"required"`
	ReviewDate time.Time `json:"review_date" validate:"required"`
	Rating     int       `json:"rating" validate:"required,min=1,max=5"`
	Comments   string    `json:"comments" validate:"required"`
	Goals      string    `json:"goals"`
}

// PerformanceReviewUpdate represents data for updating a performance review
type PerformanceReviewUpdate struct {
	Rating   int    `json:"rating" validate:"min=1,max=5"`
	Comments string `json:"comments"`
	Goals    string `json:"goals"`
}

// PerformanceReviewResponse represents performance review data returned in API responses
type PerformanceReviewResponse struct {
	ID         uuid.UUID `json:"id"`
	EmployeeID uuid.UUID `json:"employee_id"`
	ReviewerID uuid.UUID `json:"reviewer_id"`
	ReviewDate time.Time `json:"review_date"`
	Rating     int       `json:"rating"`
	Comments   string    `json:"comments"`
	Goals      string    `json:"goals"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName specifies the table name for PerformanceReview model
func (PerformanceReview) TableName() string {
	return "performance_reviews"
}
