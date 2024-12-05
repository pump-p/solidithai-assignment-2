package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pump-p/solidithai-assignment-2/backend/services"
)

func GetLogs(c *gin.Context) {
	query := c.Query("q")
	logs, err := services.QueryLogs(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}
