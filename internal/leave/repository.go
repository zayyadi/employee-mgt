package leave

import (
	"employee-management/internal/database"
	"employee-management/internal/models"
	"errors"

	"github.com/google/uuid"
)

// Repository defines the interface for leave data operations
type Repository interface {
	// Leave Type methods
	CreateLeaveType(leaveTypeData *models.LeaveTypeCreate) (*models.LeaveType, error)
	GetLeaveTypeByID(id uuid.UUID) (*models.LeaveType, error)
	ListLeaveTypes() ([]models.LeaveType, error)
	UpdateLeaveType(id uuid.UUID, leaveTypeData *models.LeaveTypeUpdate) (*models.LeaveType, error)
	DeleteLeaveType(id uuid.UUID) error

	// Leave Request methods
	CreateLeaveRequest(leaveRequestData *models.LeaveRequestCreate) (*models.LeaveRequest, error)
	GetLeaveRequestByID(id uuid.UUID) (*models.LeaveRequest, error)
	ListLeaveRequests() ([]models.LeaveRequest, error)
	UpdateLeaveRequestStatus(id uuid.UUID, status string, approvedBy *uuid.UUID) (*models.LeaveRequest, error)
}

type repository struct {
	db *database.DB
}

// NewRepository creates a new leave repository
func NewRepository(db *database.DB) Repository {
	return &repository{db}
}

// CreateLeaveType creates a new leave type
func (r *repository) CreateLeaveType(leaveTypeData *models.LeaveTypeCreate) (*models.LeaveType, error) {
	var leaveType models.LeaveType
	query := `INSERT INTO leave_types (name, description, max_days_per_year, is_accrued)
			  VALUES ($1, $2, $3, $4)
			  RETURNING id, name, description, max_days_per_year, is_accrued, created_at, updated_at`
	err := r.db.QueryRow(query, leaveTypeData.Name, leaveTypeData.Description, leaveTypeData.MaxDaysPerYear, leaveTypeData.IsAccrued).Scan(
		&leaveType.ID, &leaveType.Name, &leaveType.Description, &leaveType.MaxDaysPerYear, &leaveType.IsAccrued, &leaveType.CreatedAt, &leaveType.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &leaveType, nil
}

// GetLeaveTypeByID retrieves a leave type by ID
func (r *repository) GetLeaveTypeByID(id uuid.UUID) (*models.LeaveType, error) {
	var leaveType models.LeaveType
	query := `SELECT id, name, description, max_days_per_year, is_accrued, created_at, updated_at
			  FROM leave_types WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&leaveType.ID, &leaveType.Name, &leaveType.Description, &leaveType.MaxDaysPerYear, &leaveType.IsAccrued, &leaveType.CreatedAt, &leaveType.UpdatedAt,
	)
	if err != nil {
		return nil, errors.New("leave type not found")
	}
	return &leaveType, nil
}

// ListLeaveTypes retrieves all leave types
func (r *repository) ListLeaveTypes() ([]models.LeaveType, error) {
	var leaveTypes []models.LeaveType
	query := `SELECT id, name, description, max_days_per_year, is_accrued, created_at, updated_at
			  FROM leave_types`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var leaveType models.LeaveType
		if err := rows.Scan(&leaveType.ID, &leaveType.Name, &leaveType.Description, &leaveType.MaxDaysPerYear, &leaveType.IsAccrued, &leaveType.CreatedAt, &leaveType.UpdatedAt); err != nil {
			return nil, err
		}
		leaveTypes = append(leaveTypes, leaveType)
	}
	return leaveTypes, nil
}

// UpdateLeaveType updates a leave type
func (r *repository) UpdateLeaveType(id uuid.UUID, leaveTypeData *models.LeaveTypeUpdate) (*models.LeaveType, error) {
	var leaveType models.LeaveType
	query := `UPDATE leave_types
			  SET name = $1, description = $2, max_days_per_year = $3, is_accrued = $4, updated_at = NOW()
			  WHERE id = $5
			  RETURNING id, name, description, max_days_per_year, is_accrued, created_at, updated_at`
	err := r.db.QueryRow(query, leaveTypeData.Name, leaveTypeData.Description, leaveTypeData.MaxDaysPerYear, leaveTypeData.IsAccrued, id).Scan(
		&leaveType.ID, &leaveType.Name, &leaveType.Description, &leaveType.MaxDaysPerYear, &leaveType.IsAccrued, &leaveType.CreatedAt, &leaveType.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &leaveType, nil
}

// DeleteLeaveType deletes a leave type
func (r *repository) DeleteLeaveType(id uuid.UUID) error {
	result, err := r.db.Exec("DELETE FROM leave_types WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("leave type not found")
	}
	return nil
}

// CreateLeaveRequest creates a new leave request
func (r *repository) CreateLeaveRequest(leaveRequestData *models.LeaveRequestCreate) (*models.LeaveRequest, error) {
	var leaveRequest models.LeaveRequest
	query := `INSERT INTO leave_requests (employee_id, leave_type_id, start_date, end_date, reason)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, employee_id, leave_type_id, start_date, end_date, reason, status, approved_by, approved_at, created_at, updated_at`
	err := r.db.QueryRow(query, leaveRequestData.EmployeeID, leaveRequestData.LeaveTypeID, leaveRequestData.StartDate, leaveRequestData.EndDate, leaveRequestData.Reason).Scan(
		&leaveRequest.ID, &leaveRequest.EmployeeID, &leaveRequest.LeaveTypeID, &leaveRequest.StartDate, &leaveRequest.EndDate, &leaveRequest.Reason, &leaveRequest.Status, &leaveRequest.ApprovedBy, &leaveRequest.ApprovedAt, &leaveRequest.CreatedAt, &leaveRequest.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &leaveRequest, nil
}

// GetLeaveRequestByID retrieves a leave request by ID
func (r *repository) GetLeaveRequestByID(id uuid.UUID) (*models.LeaveRequest, error) {
	var leaveRequest models.LeaveRequest
	query := `SELECT id, employee_id, leave_type_id, start_date, end_date, reason, status, approved_by, approved_at, created_at, updated_at
			  FROM leave_requests WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&leaveRequest.ID, &leaveRequest.EmployeeID, &leaveRequest.LeaveTypeID, &leaveRequest.StartDate, &leaveRequest.EndDate, &leaveRequest.Reason, &leaveRequest.Status, &leaveRequest.ApprovedBy, &leaveRequest.ApprovedAt, &leaveRequest.CreatedAt, &leaveRequest.UpdatedAt,
	)
	if err != nil {
		return nil, errors.New("leave request not found")
	}
	return &leaveRequest, nil
}

// ListLeaveRequests retrieves all leave requests
func (r *repository) ListLeaveRequests() ([]models.LeaveRequest, error) {
	var leaveRequests []models.LeaveRequest
	query := `SELECT id, employee_id, leave_type_id, start_date, end_date, reason, status, approved_by, approved_at, created_at, updated_at
			  FROM leave_requests`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var leaveRequest models.LeaveRequest
		if err := rows.Scan(&leaveRequest.ID, &leaveRequest.EmployeeID, &leaveRequest.LeaveTypeID, &leaveRequest.StartDate, &leaveRequest.EndDate, &leaveRequest.Reason, &leaveRequest.Status, &leaveRequest.ApprovedBy, &leaveRequest.ApprovedAt, &leaveRequest.CreatedAt, &leaveRequest.UpdatedAt); err != nil {
			return nil, err
		}
		leaveRequests = append(leaveRequests, leaveRequest)
	}
	return leaveRequests, nil
}

// UpdateLeaveRequestStatus updates the status of a leave request
func (r *repository) UpdateLeaveRequestStatus(id uuid.UUID, status string, approvedBy *uuid.UUID) (*models.LeaveRequest, error) {
	var leaveRequest models.LeaveRequest
	query := `UPDATE leave_requests
			  SET status = $1, approved_by = $2, approved_at = NOW(), updated_at = NOW()
			  WHERE id = $3
			  RETURNING id, employee_id, leave_type_id, start_date, end_date, reason, status, approved_by, approved_at, created_at, updated_at`
	err := r.db.QueryRow(query, status, approvedBy, id).Scan(
		&leaveRequest.ID, &leaveRequest.EmployeeID, &leaveRequest.LeaveTypeID, &leaveRequest.StartDate, &leaveRequest.EndDate, &leaveRequest.Reason, &leaveRequest.Status, &leaveRequest.ApprovedBy, &leaveRequest.ApprovedAt, &leaveRequest.CreatedAt, &leaveRequest.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &leaveRequest, nil
}
