package handlers

import (
	"gin-api/api/dto"
	services "gin-api/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoverHandler handles rover-related requests
type RoverHandler struct {
	service *services.RoverService
}

// NewRoverHandler creates a new RoverHandler
func NewRoverHandler() *RoverHandler {
	return &RoverHandler{service: services.NewRoverService()}
}

// MoveRover handles the move rover request
func (h *RoverHandler) MoveRover(c *gin.Context) {
	var input dto.RoverDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedPosition := h.service.MoveRover(input)
	c.JSON(http.StatusOK, gin.H{"updatedPosition": updatedPosition})
}
