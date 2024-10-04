package services

import (
	"gin-api/api/dto"
	"gin-api/api/models"
)

// RoverService handles rover movements
type RoverService struct{}

// NewRoverService creates a new RoverService
func NewRoverService() *RoverService {
	return &RoverService{}
}

// MoveRover processes movement instructions for the rover
func (s *RoverService) MoveRover(input dto.RoverDTO) string {
	// Parse initial position
	x := int(input.InitialPosition[0] - '0')
	y := int(input.InitialPosition[1] - '0')
	direction := string(input.InitialPosition[2])

	// Create a new Rover
	rover := models.NewRover(x, y, direction)

	// Process each instruction
	for _, instruction := range input.Instruction {
		switch instruction {
		case 'M':
			rover.MoveForward()
		case 'L':
			rover.TurnLeft()
		case 'R':
			rover.TurnRight()
		}
	}

	return rover.GetPosition()
}
