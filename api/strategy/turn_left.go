package strategy

import "gin-api/api/models"

// TurnLeft implements MovementStrategy for turning left
type TurnLeft struct {
	rover *models.Rover
}

// NewTurnLeft creates a new TurnLeft strategy
func NewTurnLeft(rover *models.Rover) *TurnLeft {
	return &TurnLeft{rover: rover}
}

// Execute turns the rover left
func (t *TurnLeft) Execute() {
	t.rover.TurnLeft()
}
