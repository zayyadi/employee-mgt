package employee

import (
	"employee-management/internal/models"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateEmployee(employeeData *models.EmployeeCreate) (*models.Employee, error) {

	return s.repo.CreateEmployee(employeeData)
}

func (s *Service) GetEmployeeByID(id uuid.UUID) (*models.Employee, error) {
	return s.repo.GetEmployeeByID(id)
}

func (s *Service) UpdateEmployee(id uuid.UUID, employeeData *models.EmployeeUpdate) (*models.Employee, error) {

	return s.repo.UpdateEmployee(id, employeeData)
}

func (s *Service) DeleteEmployee(id uuid.UUID) error {
	return s.repo.DeleteEmployee(id)
}

func (s *Service) ListEmployees() ([]models.Employee, error) {
	return s.repo.ListEmployees()
}
