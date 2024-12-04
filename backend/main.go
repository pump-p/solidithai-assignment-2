package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pump-p/solidithai-assignment-2/backend/database"
	"github.com/pump-p/solidithai-assignment-2/backend/routes"
)

func main() {
	// Connect to the database
	database.ConnectDatabase()

	// Set up the router
	r := gin.Default()

	// Register routes
	routes.SetupRouter(r)

	// Start the server
	r.Run(":8080")
}
