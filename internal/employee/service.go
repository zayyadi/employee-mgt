package employee

import (
	"employee-management/internal/models"

	"github.com/google/uuid"
)

// Service handles employee-related operations
type Service struct {
	repo Repository
}

// NewService creates a new employee service
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateEmployee creates a new employee
func (s *Service) CreateEmployee(employeeData *models.EmployeeCreate) (*models.Employee, error) {
	// In a real application, you might have some business logic here
	// For example, validating that the specified manager_id exists and is a manager.
	return s.repo.CreateEmployee(employeeData)
}

// GetEmployeeByID retrieves an employee by their ID
func (s *Service) GetEmployeeByID(id uuid.UUID) (*models.Employee, error) {
	return s.repo.GetEmployeeByID(id)
}

// UpdateEmployee updates an existing employee's information
func (s *Service) UpdateEmployee(id uuid.UUID, employeeData *models.EmployeeUpdate) (*models.Employee, error) {
	// Business logic can be added here, e.g., checking for valid department or position IDs.
	return s.repo.UpdateEmployee(id, employeeData)
}

// DeleteEmployee deletes an employee by their ID
func (s *Service) DeleteEmployee(id uuid.UUID) error {
	return s.repo.DeleteEmployee(id)
}

// ListEmployees retrieves a list of all employees
func (s *Service) ListEmployees() ([]models.Employee, error) {
	return s.repo.ListEmployees()
}
