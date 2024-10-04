package models

import "fmt"

type Rover struct {
	X         int
	Y         int
	Direction string
}

// NewRover creates a new rover
func NewRover(x, y int, direction string) *Rover {
	return &Rover{X: x, Y: y, Direction: direction}
}

// MoveForward moves the rover one step forward in the current direction
func (r *Rover) MoveForward() {
	switch r.Direction {
	case "N":
		r.Y++
	case "E":
		r.X++
	case "S":
		r.Y--
	case "W":
		r.X--
	}
}

// TurnLeft changes the rover's direction to the left
func (r *Rover) TurnLeft() {
	switch r.Direction {
	case "N":
		r.Direction = "W"
	case "W":
		r.Direction = "S"
	case "S":
		r.Direction = "E"
	case "E":
		r.Direction = "N"
	}
}

// TurnRight changes the rover's direction to the right
func (r *Rover) TurnRight() {
	switch r.Direction {
	case "N":
		r.Direction = "E"
	case "E":
		r.Direction = "S"
	case "S":
		r.Direction = "W"
	case "W":
		r.Direction = "N"
	}
}

// GetPosition returns the current position of the rover
func (r *Rover) GetPosition() string {
	// Use fmt.Sprintf to format the string correctly
	return fmt.Sprintf("%d %d %s", r.X, r.Y, r.Direction)
}
