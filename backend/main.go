package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pump-p/solidithai-assignment-2/backend/database"
	"github.com/pump-p/solidithai-assignment-2/backend/models"
	"github.com/pump-p/solidithai-assignment-2/backend/routes"
	"github.com/pump-p/solidithai-assignment-2/backend/utils"
)

func main() {
	// Connect to the database
	database.ConnectDatabase()

	// Migrate the models
	database.DB.AutoMigrate(&models.User{})

	// Set up the router
	r := gin.Default()

	// Register routes
	routes.SetupRouter(r)

	// Run Gin server in a goroutine
	go func() {
		r.Run(":8080")
	}()

	// Set up WebSocket server
	http.HandleFunc("/ws", utils.HandleWebSocketConnections)

	// Run WebSocket server
	http.ListenAndServe(":8081", nil)
}
