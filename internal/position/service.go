package position

import (
	"employee-management/internal/database"
	"employee-management/internal/models"
	"errors"

	"github.com/google/uuid"
)

// Service handles position-related operations
type Service struct {
	db *database.DB
}

// NewService creates a new position service
func NewService(db *database.DB) *Service {
	return &Service{
		db: db,
	}
}

// CreatePosition creates a new position
func (s *Service) CreatePosition(positionData *models.PositionCreate) (*models.Position, error) {
	var position models.Position
	query := `
		INSERT INTO positions (title, department_id, description, requirements, salary_range_min, salary_range_max)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, department_id, description, requirements, salary_range_min, salary_range_max, created_at, updated_at
	`
	err := s.db.QueryRow(query,
		positionData.Title, positionData.DepartmentID, positionData.Description, positionData.Requirements, positionData.SalaryRangeMin, positionData.SalaryRangeMax,
	).Scan(
		&position.ID, &position.Title, &position.DepartmentID, &position.Description, &position.Requirements, &position.SalaryRangeMin, &position.SalaryRangeMax, &position.CreatedAt, &position.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &position, nil
}

// GetPositionByID retrieves a position by its ID
func (s *Service) GetPositionByID(id uuid.UUID) (*models.Position, error) {
	var position models.Position
	query := `
		SELECT id, title, department_id, description, requirements, salary_range_min, salary_range_max, created_at, updated_at
		FROM positions WHERE id = $1
	`
	err := s.db.QueryRow(query, id).Scan(
		&position.ID, &position.Title, &position.DepartmentID, &position.Description, &position.Requirements, &position.SalaryRangeMin, &position.SalaryRangeMax, &position.CreatedAt, &position.UpdatedAt,
	)

	if err != nil {
		return nil, errors.New("position not found")
	}

	return &position, nil
}

// UpdatePosition updates an existing position's information
func (s *Service) UpdatePosition(id uuid.UUID, positionData *models.PositionUpdate) (*models.Position, error) {
	var position models.Position
	query := `
		UPDATE positions
		SET title = $1, description = $2, requirements = $3, salary_range_min = $4, salary_range_max = $5, updated_at = NOW()
		WHERE id = $6
		RETURNING id, title, department_id, description, requirements, salary_range_min, salary_range_max, created_at, updated_at
	`
	err := s.db.QueryRow(query,
		positionData.Title, positionData.Description, positionData.Requirements, positionData.SalaryRangeMin, positionData.SalaryRangeMax, id,
	).Scan(
		&position.ID, &position.Title, &position.DepartmentID, &position.Description, &position.Requirements, &position.SalaryRangeMin, &position.SalaryRangeMax, &position.CreatedAt, &position.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &position, nil
}

// DeletePosition deletes a position by its ID
func (s *Service) DeletePosition(id uuid.UUID) error {
	result, err := s.db.Exec("DELETE FROM positions WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("position not found")
	}

	return nil
}

// ListPositions retrieves a list of all positions
func (s *Service) ListPositions() ([]models.Position, error) {
	var positions []models.Position
	query := `
		SELECT id, title, department_id, description, requirements, salary_range_min, salary_range_max, created_at, updated_at
		FROM positions
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var position models.Position
		err := rows.Scan(
			&position.ID, &position.Title, &position.DepartmentID, &position.Description, &position.Requirements, &position.SalaryRangeMin, &position.SalaryRangeMax, &position.CreatedAt, &position.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}

	return positions, nil
}
