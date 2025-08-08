package document

import (
	"employee-management/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Handler handles HTTP requests for documents
type Handler struct {
	service *Service
}

// NewHandler creates a new document handler
func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

// UploadDocument handles document uploads
func (h *Handler) UploadDocument(c *gin.Context) {
	// Parse multipart form
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10 MB
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing form"})
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	employeeID, err := uuid.Parse(c.PostForm("employeeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employeeId"})
		return
	}

	uploadedBy, err := uuid.Parse(c.PostForm("uploadedBy"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid uploadedBy"})
		return
	}

	docCreate := &models.DocumentCreate{
		EmployeeID:  employeeID,
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Category:    c.PostForm("category"),
		UploadedBy:  uploadedBy,
	}

	doc, err := h.service.UploadDocument(docCreate, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload document", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, doc)
}

// GetDocument retrieves a single document
func (h *Handler) GetDocument(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	doc, err := h.service.GetDocumentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	c.JSON(http.StatusOK, doc)
}

// ListDocuments retrieves documents for an employee
func (h *Handler) ListDocuments(c *gin.Context) {
	employeeID, err := uuid.Parse(c.Query("employeeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employeeId query parameter"})
		return
	}

	docs, err := h.service.ListDocumentsByEmployeeID(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list documents"})
		return
	}

	c.JSON(http.StatusOK, docs)
}

// DeleteDocument deletes a document
func (h *Handler) DeleteDocument(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.service.DeleteDocument(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete document"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
