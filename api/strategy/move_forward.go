package strategy

import "gin-api/api/models"

// MoveForward implements MovementStrategy for moving forward
type MoveForward struct {
	rover *models.Rover
}

// NewMoveForward creates a new MoveForward strategy
func NewMoveForward(rover *models.Rover) *MoveForward {
	return &MoveForward{rover: rover}
}

// Execute moves the rover forward
func (m *MoveForward) Execute() {
	m.rover.MoveForward()
}
