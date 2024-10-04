package main

import (
	"gin-api/api/handlers"
	"gin-api/api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	roverHandler := handlers.NewRoverHandler()
	routers.SetupRoutes(router, roverHandler)

	router.Run(":8080") // Start the server on port 8080
}
