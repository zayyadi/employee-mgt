package employee

import (
	"employee-management/internal/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
func (s *Service) CreateEmployee(logger *logrus.Entry, employeeData *models.EmployeeCreate) (*models.Employee, error) {
	logger.Info("Creating a new employee")
	return s.repo.CreateEmployee(logger, employeeData)
}

// GetEmployeeByID retrieves an employee by their ID
func (s *Service) GetEmployeeByID(logger *logrus.Entry, id uuid.UUID) (*models.Employee, error) {
	logger.WithField("employeeID", id).Info("Getting employee by ID")
	return s.repo.GetEmployeeByID(logger, id)
}

// UpdateEmployee updates an existing employee's information
func (s *Service) UpdateEmployee(logger *logrus.Entry, id uuid.UUID, employeeData *models.EmployeeUpdate) (*models.Employee, error) {
	logger.WithField("employeeID", id).Info("Updating employee")
	return s.repo.UpdateEmployee(logger, id, employeeData)
}

// DeleteEmployee deletes an employee by their ID
func (s *Service) DeleteEmployee(logger *logrus.Entry, id uuid.UUID) error {
	logger.WithField("employeeID", id).Info("Deleting employee")
	return s.repo.DeleteEmployee(logger, id)
}

// ListEmployees retrieves a list of all employees
func (s *Service) ListEmployees(logger *logrus.Entry) ([]models.Employee, error) {
	logger.Info("Listing all employees")
	return s.repo.ListEmployees(logger)
}
