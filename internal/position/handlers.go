package position

import (
	"employee-management/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Handler handles position-related HTTP requests
type Handler struct {
	service *Service
}

// NewHandler creates a new position handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreatePosition handles the creation of a new position
// @Summary Create a new position
// @Description Create a new position with the provided data
// @Tags Positions
// @Accept json
// @Produce json
// @Param position body models.PositionCreate true "Position data"
// @Success 201 {object} models.Position
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /positions [post]
func (h *Handler) CreatePosition(c *gin.Context) {
	var positionData models.PositionCreate
	if err := c.ShouldBindJSON(&positionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	position, err := h.service.CreatePosition(&positionData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, position)
}

// GetPosition handles retrieving a position by its ID
// @Summary Get a position by ID
// @Description Get a position's details by its ID
// @Tags Positions
// @Produce json
// @Param id path string true "Position ID"
// @Success 200 {object} models.Position
// @Failure 404 {object} map[string]string
// @Router /positions/{id} [get]
func (h *Handler) GetPosition(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid position ID"})
		return
	}

	position, err := h.service.GetPositionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, position)
}

// UpdatePosition handles updating an existing position
// @Summary Update a position
// @Description Update an existing position's information
// @Tags Positions
// @Accept json
// @Produce json
// @Param id path string true "Position ID"
// @Param position body models.PositionUpdate true "Position data"
// @Success 200 {object} models.Position
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /positions/{id} [put]
func (h *Handler) UpdatePosition(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid position ID"})
		return
	}

	var positionData models.PositionUpdate
	if err := c.ShouldBindJSON(&positionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	position, err := h.service.UpdatePosition(id, &positionData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, position)
}

// DeletePosition handles deleting a position by its ID
// @Summary Delete a position
// @Description Delete a position by its ID
// @Tags Positions
// @Produce json
// @Param id path string true "Position ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /positions/{id} [delete]
func (h *Handler) DeletePosition(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid position ID"})
		return
	}

	err = h.service.DeletePosition(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListPositions handles listing all positions
// @Summary List all positions
// @Description Get a list of all positions
// @Tags Positions
// @Produce json
// @Success 200 {array} models.Position
// @Failure 500 {object} map[string]string
// @Router /positions [get]
func (h *Handler) ListPositions(c *gin.Context) {
	positions, err := h.service.ListPositions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, positions)
}
