package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pump-p/solidithai-assignment-2/backend/database"
	"github.com/pump-p/solidithai-assignment-2/backend/routes"
)

func main() {
	// Connect to the database
	database.ConnectDatabase()

	// Set up the router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,                                                // Allow cookies/authentication
	}))

	// Register routes
	routes.SetupRouter(r)

	// Start the server
	r.Run(":5555")
}
