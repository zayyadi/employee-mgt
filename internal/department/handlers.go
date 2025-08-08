package department

import (
	"employee-management/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Handler handles department-related HTTP requests
type Handler struct {
	service *Service
}

// NewHandler creates a new department handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateDepartment handles the creation of a new department
// @Summary Create a new department
// @Description Create a new department with the provided data
// @Tags Departments
// @Accept json
// @Produce json
// @Param department body models.DepartmentCreate true "Department data"
// @Success 201 {object} models.Department
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /departments [post]
func (h *Handler) CreateDepartment(c *gin.Context) {
	var departmentData models.DepartmentCreate
	if err := c.ShouldBindJSON(&departmentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department, err := h.service.CreateDepartment(&departmentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, department)
}

// GetDepartment handles retrieving a department by its ID
// @Summary Get a department by ID
// @Description Get a department's details by its ID
// @Tags Departments
// @Produce json
// @Param id path string true "Department ID"
// @Success 200 {object} models.Department
// @Failure 404 {object} map[string]string
// @Router /departments/{id} [get]
func (h *Handler) GetDepartment(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	department, err := h.service.GetDepartmentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, department)
}

// UpdateDepartment handles updating an existing department
// @Summary Update a department
// @Description Update an existing department's information
// @Tags Departments
// @Accept json
// @Produce json
// @Param id path string true "Department ID"
// @Param department body models.DepartmentUpdate true "Department data"
// @Success 200 {object} models.Department
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /departments/{id} [put]
func (h *Handler) UpdateDepartment(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	var departmentData models.DepartmentUpdate
	if err := c.ShouldBindJSON(&departmentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department, err := h.service.UpdateDepartment(id, &departmentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, department)
}

// DeleteDepartment handles deleting a department by its ID
// @Summary Delete a department
// @Description Delete a department by its ID
// @Tags Departments
// @Produce json
// @Param id path string true "Department ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /departments/{id} [delete]
func (h *Handler) DeleteDepartment(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	err = h.service.DeleteDepartment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListDepartments handles listing all departments
// @Summary List all departments
// @Description Get a list of all departments
// @Tags Departments
// @Produce json
// @Success 200 {array} models.Department
// @Failure 500 {object} map[string]string
// @Router /departments [get]
func (h *Handler) ListDepartments(c *gin.Context) {
	departments, err := h.service.ListDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, departments)
}
