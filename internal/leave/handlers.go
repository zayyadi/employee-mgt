package leave

import (
	"employee-management/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Handler handles HTTP requests for leave management
type Handler struct {
	service *Service
}

// NewHandler creates a new leave handler
func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

// Leave Type Handlers

func (h *Handler) CreateLeaveType(c *gin.Context) {
	var input models.LeaveTypeCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leaveType, err := h.service.CreateLeaveType(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create leave type"})
		return
	}

	c.JSON(http.StatusCreated, leaveType)
}

func (h *Handler) GetLeaveType(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	leaveType, err := h.service.GetLeaveTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Leave type not found"})
		return
	}

	c.JSON(http.StatusOK, leaveType)
}

func (h *Handler) ListLeaveTypes(c *gin.Context) {
	leaveTypes, err := h.service.ListLeaveTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list leave types"})
		return
	}

	c.JSON(http.StatusOK, leaveTypes)
}

func (h *Handler) UpdateLeaveType(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input models.LeaveTypeUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leaveType, err := h.service.UpdateLeaveType(id, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update leave type"})
		return
	}

	c.JSON(http.StatusOK, leaveType)
}

func (h *Handler) DeleteLeaveType(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.service.DeleteLeaveType(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete leave type"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Leave Request Handlers

func (h *Handler) CreateLeaveRequest(c *gin.Context) {
	var input models.LeaveRequestCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real app, you would get the employee ID from the authenticated user
	// For now, we'll trust the input.

	leaveRequest, err := h.service.CreateLeaveRequest(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create leave request"})
		return
	}

	c.JSON(http.StatusCreated, leaveRequest)
}

func (h *Handler) GetLeaveRequest(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	leaveRequest, err := h.service.GetLeaveRequestByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Leave request not found"})
		return
	}

	c.JSON(http.StatusOK, leaveRequest)
}

func (h *Handler) ListLeaveRequests(c *gin.Context) {
	leaveRequests, err := h.service.ListLeaveRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list leave requests"})
		return
	}

	c.JSON(http.StatusOK, leaveRequests)
}

func (h *Handler) ApproveLeaveRequest(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// In a real app, you would get the approver's user ID from the authenticated user (e.g. a manager)
	// For simplicity, we'll use a placeholder UUID.
	approverID := uuid.New() // Placeholder

	leaveRequest, err := h.service.ApproveLeaveRequest(id, approverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leaveRequest)
}

func (h *Handler) RejectLeaveRequest(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Similar to approve, get rejector's ID from auth context.
	rejectorID := uuid.New() // Placeholder

	leaveRequest, err := h.service.RejectLeaveRequest(id, rejectorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leaveRequest)
}
