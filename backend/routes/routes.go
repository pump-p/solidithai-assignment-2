package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pump-p/solidithai-assignment-2/backend/controllers"
	"github.com/pump-p/solidithai-assignment-2/backend/middlewares"
)

func SetupRouter(router *gin.Engine) {

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accpet"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/signup", controllers.Signup)
		authRoutes.POST("/login", controllers.Login)
	}

	userRoutes := router.Group("/users").Use(middlewares.AuthMiddleware())
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUserByID)
		userRoutes.POST("", controllers.CreateUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	// Protected routes for log querying
	logRoutes := router.Group("/logs")
	logRoutes.Use(middlewares.AuthMiddleware())
	{
		logRoutes.GET("", controllers.GetLogs)
	}
}
