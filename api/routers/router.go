package routers

import (
	"gin-api/api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes for the API
func SetupRoutes(router *gin.Engine, roverHandler *handlers.RoverHandler) {
	router.POST("/move", roverHandler.MoveRover) // Define the route for moving the rover
}
