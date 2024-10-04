// api/forms/rover_form.go
package forms

import "gin-api/api/dto"

// RoverForm is a struct to handle rover input form data
type RoverForm struct {
	DTO dto.RoverDTO
}

// Validate checks if the RoverForm is valid
func (f *RoverForm) Validate() error {
	// Add validation logic here if needed
	return nil
}
