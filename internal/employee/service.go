package employee

import (
	"employee-management/internal/database"
	"employee-management/internal/models"
	"errors"

	"github.com/google/uuid"
)

// Service handles employee-related operations
type Service struct {
	db *database.DB
}

// NewService creates a new employee service
func NewService(db *database.DB) *Service {
	return &Service{
		db: db,
	}
}

// CreateEmployee creates a new employee
func (s *Service) CreateEmployee(employeeData *models.EmployeeCreate) (*models.Employee, error) {
	var employee models.Employee
	query := `
		INSERT INTO employees (user_id, employee_id, first_name, last_name, date_of_birth, gender, marital_status, phone_number, email, address, emergency_contact_name, emergency_contact_phone, department_id, position_id, hire_date, employment_status, manager_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		RETURNING id, user_id, employee_id, first_name, last_name, date_of_birth, gender, marital_status, phone_number, email, address, emergency_contact_name, emergency_contact_phone, department_id, position_id, hire_date, employment_status, manager_id, created_at, updated_at
	`
	err := s.db.QueryRow(query,
		employeeData.UserID, employeeData.EmployeeID, employeeData.FirstName, employeeData.LastName, employeeData.DateOfBirth, employeeData.Gender, employeeData.MaritalStatus, employeeData.PhoneNumber, employeeData.Email, employeeData.Address, employeeData.EmergencyContactName, employeeData.EmergencyContactPhone, employeeData.DepartmentID, employeeData.PositionID, employeeData.HireDate, employeeData.EmploymentStatus, employeeData.ManagerID,
	).Scan(
		&employee.ID, &employee.UserID, &employee.EmployeeID, &employee.FirstName, &employee.LastName, &employee.DateOfBirth, &employee.Gender, &employee.MaritalStatus, &employee.PhoneNumber, &employee.Email, &employee.Address, &employee.EmergencyContactName, &employee.EmergencyContactPhone, &employee.DepartmentID, &employee.PositionID, &employee.HireDate, &employee.EmploymentStatus, &employee.ManagerID, &employee.CreatedAt, &employee.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

// GetEmployeeByID retrieves an employee by their ID
func (s *Service) GetEmployeeByID(id uuid.UUID) (*models.Employee, error) {
	var employee models.Employee
	query := `
		SELECT id, user_id, employee_id, first_name, last_name, date_of_birth, gender, marital_status, phone_number, email, address, emergency_contact_name, emergency_contact_phone, department_id, position_id, hire_date, employment_status, manager_id, created_at, updated_at
		FROM employees WHERE id = $1
	`
	err := s.db.QueryRow(query, id).Scan(
		&employee.ID, &employee.UserID, &employee.EmployeeID, &employee.FirstName, &employee.LastName, &employee.DateOfBirth, &employee.Gender, &employee.MaritalStatus, &employee.PhoneNumber, &employee.Email, &employee.Address, &employee.EmergencyContactName, &employee.EmergencyContactPhone, &employee.DepartmentID, &employee.PositionID, &employee.HireDate, &employee.EmploymentStatus, &employee.ManagerID, &employee.CreatedAt, &employee.UpdatedAt,
	)

	if err != nil {
		return nil, errors.New("employee not found")
	}

	return &employee, nil
}

// UpdateEmployee updates an existing employee's information
func (s *Service) UpdateEmployee(id uuid.UUID, employeeData *models.EmployeeUpdate) (*models.Employee, error) {
	var employee models.Employee
	query := `
		UPDATE employees
		SET first_name = $1, last_name = $2, date_of_birth = $3, gender = $4, marital_status = $5, phone_number = $6, email = $7, address = $8, emergency_contact_name = $9, emergency_contact_phone = $10, department_id = $11, position_id = $12, hire_date = $13, employment_status = $14, manager_id = $15, updated_at = NOW()
		WHERE id = $16
		RETURNING id, user_id, employee_id, first_name, last_name, date_of_birth, gender, marital_status, phone_number, email, address, emergency_contact_name, emergency_contact_phone, department_id, position_id, hire_date, employment_status, manager_id, created_at, updated_at
	`
	err := s.db.QueryRow(query,
		employeeData.FirstName, employeeData.LastName, employeeData.DateOfBirth, employeeData.Gender, employeeData.MaritalStatus, employeeData.PhoneNumber, employeeData.Email, employeeData.Address, employeeData.EmergencyContactName, employeeData.EmergencyContactPhone, employeeData.DepartmentID, employeeData.PositionID, employeeData.HireDate, employeeData.EmploymentStatus, employeeData.ManagerID, id,
	).Scan(
		&employee.ID, &employee.UserID, &employee.EmployeeID, &employee.FirstName, &employee.LastName, &employee.DateOfBirth, &employee.Gender, &employee.MaritalStatus, &employee.PhoneNumber, &employee.Email, &employee.Address, &employee.EmergencyContactName, &employee.EmergencyContactPhone, &employee.DepartmentID, &employee.PositionID, &employee.HireDate, &employee.EmploymentStatus, &employee.ManagerID, &employee.CreatedAt, &employee.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

// DeleteEmployee deletes an employee by their ID
func (s *Service) DeleteEmployee(id uuid.UUID) error {
	result, err := s.db.Exec("DELETE FROM employees WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("employee not found")
	}

	return nil
}

// ListEmployees retrieves a list of all employees
func (s *Service) ListEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	query := `
		SELECT id, user_id, employee_id, first_name, last_name, date_of_birth, gender, marital_status, phone_number, email, address, emergency_contact_name, emergency_contact_phone, department_id, position_id, hire_date, employment_status, manager_id, created_at, updated_at
		FROM employees
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		err := rows.Scan(
			&employee.ID, &employee.UserID, &employee.EmployeeID, &employee.FirstName, &employee.LastName, &employee.DateOfBirth, &employee.Gender, &employee.MaritalStatus, &employee.PhoneNumber, &employee.Email, &employee.Address, &employee.EmergencyContactName, &employee.EmergencyContactPhone, &employee.DepartmentID, &employee.PositionID, &employee.HireDate, &employee.EmploymentStatus, &employee.ManagerID, &employee.CreatedAt, &employee.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
