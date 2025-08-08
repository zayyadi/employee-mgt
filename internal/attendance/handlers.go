package attendance

import (
	"employee-management/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Handler handles attendance-related HTTP requests
type Handler struct {
	service *Service
}

// NewHandler creates a new attendance handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateAttendance handles the creation of a new attendance record
// @Summary Create a new attendance record
// @Description Create a new attendance record with the provided data
// @Tags Attendance
// @Accept json
// @Produce json
// @Param attendance body models.AttendanceCreate true "Attendance data"
// @Success 201 {object} models.Attendance
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance [post]
func (h *Handler) CreateAttendance(c *gin.Context) {
	var attendanceData models.AttendanceCreate
	if err := c.ShouldBindJSON(&attendanceData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendance, err := h.service.CreateAttendance(&attendanceData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, attendance)
}

// GetAttendance handles retrieving an attendance record by its ID
// @Summary Get an attendance record by ID
// @Description Get an attendance record's details by its ID
// @Tags Attendance
// @Produce json
// @Param id path string true "Attendance ID"
// @Success 200 {object} models.Attendance
// @Failure 404 {object} map[string]string
// @Router /attendance/{id} [get]
func (h *Handler) GetAttendance(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attendance ID"})
		return
	}

	attendance, err := h.service.GetAttendanceByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

// UpdateAttendance handles updating an existing attendance record
// @Summary Update an attendance record
// @Description Update an existing attendance record's information
// @Tags Attendance
// @Accept json
// @Produce json
// @Param id path string true "Attendance ID"
// @Param attendance body models.AttendanceUpdate true "Attendance data"
// @Success 200 {object} models.Attendance
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/{id} [put]
func (h *Handler) UpdateAttendance(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attendance ID"})
		return
	}

	var attendanceData models.AttendanceUpdate
	if err := c.ShouldBindJSON(&attendanceData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendance, err := h.service.UpdateAttendance(id, &attendanceData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

// DeleteAttendance handles deleting an attendance record by its ID
// @Summary Delete an attendance record
// @Description Delete an attendance record by its ID
// @Tags Attendance
// @Produce json
// @Param id path string true "Attendance ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /attendance/{id} [delete]
func (h *Handler) DeleteAttendance(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attendance ID"})
		return
	}

	err = h.service.DeleteAttendance(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListAttendance handles listing all attendance records
// @Summary List all attendance records
// @Description Get a list of all attendance records
// @Tags Attendance
// @Produce json
// @Success 200 {array} models.Attendance
// @Failure 500 {object} map[string]string
// @Router /attendance [get]
func (h *Handler) ListAttendance(c *gin.Context) {
	attendances, err := h.service.ListAttendance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendances)
}

// CheckIn handles employee check-in
// @Summary Employee check-in
// @Description Record an employee's check-in time
// @Tags Attendance
// @Accept json
// @Produce json
// @Param checkin body map[string]string true "Employee ID"
// @Success 201 {object} models.Attendance
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/check-in [post]
func (h *Handler) CheckIn(c *gin.Context) {
	var requestData struct {
		EmployeeID string `json:"employee_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employeeID, err := uuid.Parse(requestData.EmployeeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	attendance, err := h.service.CheckIn(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, attendance)
}

// CheckOut handles employee check-out
// @Summary Employee check-out
// @Description Record an employee's check-out time
// @Tags Attendance
// @Accept json
// @Produce json
// @Param checkout body map[string]string true "Employee ID"
// @Success 200 {object} models.Attendance
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/check-out [post]
func (h *Handler) CheckOut(c *gin.Context) {
	var requestData struct {
		EmployeeID string `json:"employee_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employeeID, err := uuid.Parse(requestData.EmployeeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	attendance, err := h.service.CheckOut(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}
