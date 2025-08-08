package department

import (
	"employee-management/internal/database"
	"employee-management/internal/models"
	"errors"

	"github.com/google/uuid"
)

// Service handles department-related operations
type Service struct {
	db *database.DB
}

// NewService creates a new department service
func NewService(db *database.DB) *Service {
	return &Service{
		db: db,
	}
}

// CreateDepartment creates a new department
func (s *Service) CreateDepartment(departmentData *models.DepartmentCreate) (*models.Department, error) {
	var department models.Department
	query := `
		INSERT INTO departments (name, description, manager_id)
		VALUES ($1, $2, $3)
		RETURNING id, name, description, manager_id, created_at, updated_at
	`
	err := s.db.QueryRow(query,
		departmentData.Name, departmentData.Description, departmentData.ManagerID,
	).Scan(
		&department.ID, &department.Name, &department.Description, &department.ManagerID, &department.CreatedAt, &department.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &department, nil
}

// GetDepartmentByID retrieves a department by its ID
func (s *Service) GetDepartmentByID(id uuid.UUID) (*models.Department, error) {
	var department models.Department
	query := `
		SELECT id, name, description, manager_id, created_at, updated_at
		FROM departments WHERE id = $1
	`
	err := s.db.QueryRow(query, id).Scan(
		&department.ID, &department.Name, &department.Description, &department.ManagerID, &department.CreatedAt, &department.UpdatedAt,
	)

	if err != nil {
		return nil, errors.New("department not found")
	}

	return &department, nil
}

// UpdateDepartment updates an existing department's information
func (s *Service) UpdateDepartment(id uuid.UUID, departmentData *models.DepartmentUpdate) (*models.Department, error) {
	var department models.Department
	query := `
		UPDATE departments
		SET name = $1, description = $2, manager_id = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING id, name, description, manager_id, created_at, updated_at
	`
	err := s.db.QueryRow(query,
		departmentData.Name, departmentData.Description, departmentData.ManagerID, id,
	).Scan(
		&department.ID, &department.Name, &department.Description, &department.ManagerID, &department.CreatedAt, &department.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &department, nil
}

// DeleteDepartment deletes a department by its ID
func (s *Service) DeleteDepartment(id uuid.UUID) error {
	result, err := s.db.Exec("DELETE FROM departments WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("department not found")
	}

	return nil
}

// ListDepartments retrieves a list of all departments
func (s *Service) ListDepartments() ([]models.Department, error) {
	var departments []models.Department
	query := `
		SELECT id, name, description, manager_id, created_at, updated_at
		FROM departments
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var department models.Department
		err := rows.Scan(
			&department.ID, &department.Name, &department.Description, &department.ManagerID, &department.CreatedAt, &department.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}

	return departments, nil
}
