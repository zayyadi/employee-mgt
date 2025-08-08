package employee

import (
	"employee-management/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Handler handles employee-related HTTP requests
type Handler struct {
	service *Service
}

// NewHandler creates a new employee handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateEmployee handles the creation of a new employee
// @Summary Create a new employee
// @Description Create a new employee with the provided data
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body models.EmployeeCreate true "Employee data"
// @Success 201 {object} models.Employee
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees [post]
func (h *Handler) CreateEmployee(c *gin.Context) {
	var employeeData models.EmployeeCreate
	if err := c.ShouldBindJSON(&employeeData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := h.service.CreateEmployee(&employeeData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

// GetEmployee handles retrieving an employee by their ID
// @Summary Get an employee by ID
// @Description Get an employee's details by their ID
// @Tags Employees
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} models.Employee
// @Failure 404 {object} map[string]string
// @Router /employees/{id} [get]
func (h *Handler) GetEmployee(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	employee, err := h.service.GetEmployeeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// UpdateEmployee handles updating an existing employee
// @Summary Update an employee
// @Description Update an existing employee's information
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param employee body models.EmployeeUpdate true "Employee data"
// @Success 200 {object} models.Employee
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [put]
func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var employeeData models.EmployeeUpdate
	if err := c.ShouldBindJSON(&employeeData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := h.service.UpdateEmployee(id, &employeeData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// DeleteEmployee handles deleting an employee by their ID
// @Summary Delete an employee
// @Description Delete an employee by their ID
// @Tags Employees
// @Produce json
// @Param id path string true "Employee ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /employees/{id} [delete]
func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	err = h.service.DeleteEmployee(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListEmployees handles listing all employees
// @Summary List all employees
// @Description Get a list of all employees
// @Tags Employees
// @Produce json
// @Success 200 {array} models.Employee
// @Failure 500 {object} map[string]string
// @Router /employees [get]
func (h *Handler) ListEmployees(c *gin.Context) {
	employees, err := h.service.ListEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}
