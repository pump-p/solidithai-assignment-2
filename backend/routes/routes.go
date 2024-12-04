package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pump-p/solidithai-assignment-2/backend/controllers"
	"github.com/pump-p/solidithai-assignment-2/backend/middlewares"
)

func SetupRouter(router *gin.Engine) {

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/signup", controllers.Signup)
		authRoutes.POST("/login", controllers.Login)
	}

	userRoutes := router.Group("/users").Use(middlewares.AuthMiddleware())
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUserByID)
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}
}
