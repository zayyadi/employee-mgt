package document

import (
	"employee-management/internal/models"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const uploadDir = "uploads"

// Service handles document-related business logic
type Service struct {
	repo Repository
}

// NewService creates a new document service
func NewService(repo Repository) *Service {
	return &Service{repo}
}

// UploadDocument handles saving a file and creating a document record
func (s *Service) UploadDocument(data *models.DocumentCreate, fileHeader *multipart.FileHeader) (*models.Document, error) {
	// Generate a unique filename to prevent collisions
	ext := filepath.Ext(fileHeader.Filename)
	newFileName := uuid.New().String() + ext
	filePath := filepath.Join(uploadDir, newFileName)

	// Create the file
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Open the uploaded file
	src, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Copy the uploaded file to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	// Save the document record to the database
	data.FilePath = filePath
	return s.repo.CreateDocument(data)
}

// GetDocumentByID retrieves a document by its ID
func (s *Service) GetDocumentByID(id uuid.UUID) (*models.Document, error) {
	return s.repo.GetDocumentByID(id)
}

// ListDocumentsByEmployeeID retrieves all documents for a given employee
func (s *Service) ListDocumentsByEmployeeID(employeeID uuid.UUID) ([]models.Document, error) {
	return s.repo.ListDocumentsByEmployeeID(employeeID)
}

// DeleteDocument handles deleting a file and its database record
func (s *Service) DeleteDocument(id uuid.UUID) error {
	// Get document to retrieve file path
	doc, err := s.repo.GetDocumentByID(id)
	if err != nil {
		return err
	}

	// Delete the file from the filesystem
	if err := os.Remove(doc.FilePath); err != nil {
		// Log error but continue to delete the DB record
	}

	// Delete the document record from the database
	return s.repo.DeleteDocument(id)
}
