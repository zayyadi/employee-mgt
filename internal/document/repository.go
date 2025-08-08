package document

import (
	"employee-management/internal/database"
	"employee-management/internal/models"

	"github.com/google/uuid"
)

// Repository defines the interface for document data operations
type Repository interface {
	CreateDocument(data *models.DocumentCreate) (*models.Document, error)
	GetDocumentByID(id uuid.UUID) (*models.Document, error)
	ListDocumentsByEmployeeID(employeeID uuid.UUID) ([]models.Document, error)
	DeleteDocument(id uuid.UUID) error
}

type repository struct {
	db *database.DB
}

// NewRepository creates a new document repository
func NewRepository(db *database.DB) Repository {
	return &repository{db}
}

// CreateDocument creates a new document record in the database
func (r *repository) CreateDocument(data *models.DocumentCreate) (*models.Document, error) {
	var doc models.Document
	query := `INSERT INTO documents (employee_id, name, description, file_path, category, uploaded_by)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING id, employee_id, name, description, file_path, category, uploaded_by, created_at`
	err := r.db.QueryRow(query, data.EmployeeID, data.Name, data.Description, data.FilePath, data.Category, data.UploadedBy).Scan(
		&doc.ID, &doc.EmployeeID, &doc.Name, &doc.Description, &doc.FilePath, &doc.Category, &doc.UploadedBy, &doc.CreatedAt,
	)
	return &doc, err
}

// GetDocumentByID retrieves a document by its ID
func (r *repository) GetDocumentByID(id uuid.UUID) (*models.Document, error) {
	var doc models.Document
	query := `SELECT id, employee_id, name, description, file_path, category, uploaded_by, created_at
			  FROM documents WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&doc.ID, &doc.EmployeeID, &doc.Name, &doc.Description, &doc.FilePath, &doc.Category, &doc.UploadedBy, &doc.CreatedAt,
	)
	return &doc, err
}

// ListDocumentsByEmployeeID retrieves all documents for a given employee
func (r *repository) ListDocumentsByEmployeeID(employeeID uuid.UUID) ([]models.Document, error) {
	var docs []models.Document
	query := `SELECT id, employee_id, name, description, file_path, category, uploaded_by, created_at
			  FROM documents WHERE employee_id = $1`
	rows, err := r.db.Query(query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var doc models.Document
		if err := rows.Scan(&doc.ID, &doc.EmployeeID, &doc.Name, &doc.Description, &doc.FilePath, &doc.Category, &doc.UploadedBy, &doc.CreatedAt); err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

// DeleteDocument deletes a document from the database
func (r *repository) DeleteDocument(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM documents WHERE id = $1", id)
	return err
}
