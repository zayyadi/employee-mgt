package attendance

import (
	"employee-management/internal/database"
	"employee-management/internal/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Service handles attendance-related operations
type Service struct {
	db *database.DB
}

// NewService creates a new attendance service
func NewService(db *database.DB) *Service {
	return &Service{
		db: db,
	}
}

// CreateAttendance creates a new attendance record
func (s *Service) CreateAttendance(attendanceData *models.AttendanceCreate) (*models.Attendance, error) {
	var attendance models.Attendance
	query := `
		INSERT INTO attendance (employee_id, check_in_time, date, status, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, employee_id, check_in_time, check_out_time, date, status, notes, created_at
	`
	err := s.db.QueryRow(query,
		attendanceData.EmployeeID, attendanceData.CheckInTime, attendanceData.Date, attendanceData.Status, attendanceData.Notes,
	).Scan(
		&attendance.ID, &attendance.EmployeeID, &attendance.CheckInTime, &attendance.CheckOutTime, &attendance.Date, &attendance.Status, &attendance.Notes, &attendance.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

// GetAttendanceByID retrieves an attendance record by its ID
func (s *Service) GetAttendanceByID(id uuid.UUID) (*models.Attendance, error) {
	var attendance models.Attendance
	query := `
		SELECT id, employee_id, check_in_time, check_out_time, date, status, notes, created_at
		FROM attendance WHERE id = $1
	`
	err := s.db.QueryRow(query, id).Scan(
		&attendance.ID, &attendance.EmployeeID, &attendance.CheckInTime, &attendance.CheckOutTime, &attendance.Date, &attendance.Status, &attendance.Notes, &attendance.CreatedAt,
	)

	if err != nil {
		return nil, errors.New("attendance record not found")
	}

	return &attendance, nil
}

// UpdateAttendance updates an existing attendance record
func (s *Service) UpdateAttendance(id uuid.UUID, attendanceData *models.AttendanceUpdate) (*models.Attendance, error) {
	var attendance models.Attendance
	query := `
		UPDATE attendance
		SET check_out_time = $1, status = $2, notes = $3
		WHERE id = $4
		RETURNING id, employee_id, check_in_time, check_out_time, date, status, notes, created_at
	`
	err := s.db.QueryRow(query,
		attendanceData.CheckOutTime, attendanceData.Status, attendanceData.Notes, id,
	).Scan(
		&attendance.ID, &attendance.EmployeeID, &attendance.CheckInTime, &attendance.CheckOutTime, &attendance.Date, &attendance.Status, &attendance.Notes, &attendance.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

// DeleteAttendance deletes an attendance record by its ID
func (s *Service) DeleteAttendance(id uuid.UUID) error {
	result, err := s.db.Exec("DELETE FROM attendance WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("attendance record not found")
	}

	return nil
}

// ListAttendance retrieves a list of all attendance records
func (s *Service) ListAttendance() ([]models.Attendance, error) {
	var attendances []models.Attendance
	query := `
		SELECT id, employee_id, check_in_time, check_out_time, date, status, notes, created_at
		FROM attendance
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attendance models.Attendance
		err := rows.Scan(
			&attendance.ID, &attendance.EmployeeID, &attendance.CheckInTime, &attendance.CheckOutTime, &attendance.Date, &attendance.Status, &attendance.Notes, &attendance.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

// CheckIn creates a new attendance record for an employee checking in
func (s *Service) CheckIn(employeeID uuid.UUID) (*models.Attendance, error) {
	// Check if employee already has a check-in for today
	var existingAttendance models.Attendance
	query := `
		SELECT id, employee_id, check_in_time, check_out_time, date, status, notes, created_at
		FROM attendance 
		WHERE employee_id = $1 AND date = $2 AND check_out_time IS NULL
	`
	today := time.Now().Truncate(24 * time.Hour)
	err := s.db.QueryRow(query, employeeID, today).Scan(
		&existingAttendance.ID, &existingAttendance.EmployeeID, &existingAttendance.CheckInTime, &existingAttendance.CheckOutTime, &existingAttendance.Date, &existingAttendance.Status, &existingAttendance.Notes, &existingAttendance.CreatedAt,
	)

	// If there's already a check-in for today, return an error
	if err == nil {
		return nil, errors.New("employee already checked in today")
	}

	// Create a new attendance record
	var attendance models.Attendance
	query = `
		INSERT INTO attendance (employee_id, check_in_time, date, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, employee_id, check_in_time, check_out_time, date, status, notes, created_at
	`
	checkInTime := time.Now()
	status := "present"
	if checkInTime.Hour() > 9 { // Assuming 9 AM is the expected check-in time
		status = "late"
	}

	err = s.db.QueryRow(query, employeeID, checkInTime, today, status).Scan(
		&attendance.ID, &attendance.EmployeeID, &attendance.CheckInTime, &attendance.CheckOutTime, &attendance.Date, &attendance.Status, &attendance.Notes, &attendance.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

// CheckOut updates an attendance record for an employee checking out
func (s *Service) CheckOut(employeeID uuid.UUID) (*models.Attendance, error) {
	// Find the attendance record for today where check_out_time is NULL
	var attendance models.Attendance
	query := `
		SELECT id, employee_id, check_in_time, check_out_time, date, status, notes, created_at
		FROM attendance 
		WHERE employee_id = $1 AND date = $2 AND check_out_time IS NULL
	`
	today := time.Now().Truncate(24 * time.Hour)
	err := s.db.QueryRow(query, employeeID, today).Scan(
		&attendance.ID, &attendance.EmployeeID, &attendance.CheckInTime, &attendance.CheckOutTime, &attendance.Date, &attendance.Status, &attendance.Notes, &attendance.CreatedAt,
	)

	if err != nil {
		return nil, errors.New("no check-in record found for today")
	}

	// Update the attendance record with check-out time
	checkOutTime := time.Now()
	query = `
		UPDATE attendance
		SET check_out_time = $1
		WHERE id = $2
		RETURNING id, employee_id, check_in_time, check_out_time, date, status, notes, created_at
	`
	err = s.db.QueryRow(query, checkOutTime, attendance.ID).Scan(
		&attendance.ID, &attendance.EmployeeID, &attendance.CheckInTime, &attendance.CheckOutTime, &attendance.Date, &attendance.Status, &attendance.Notes, &attendance.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &attendance, nil
}
