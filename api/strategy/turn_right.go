package strategy

import "gin-api/api/models"

// TurnRight implements MovementStrategy for turning right
type TurnRight struct {
	rover *models.Rover
}

// NewTurnRight creates a new TurnRight strategy
func NewTurnRight(rover *models.Rover) *TurnRight {
	return &TurnRight{rover: rover}
}

// Execute turns the rover right
func (t *TurnRight) Execute() {
	t.rover.TurnRight()
}
