package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthHandler returns simple ok status
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
