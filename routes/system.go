package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/openearthplatforminitiative/agriculture-api/handlers"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}

// Ready godoc
// @Summary Check if the service is ready
// @Description Performs health checks on various endpoints and returns their status.
// @Tags system
// @Produce json
// @Success 200 {object} models.HealthCheckResultSummary
// @Router /ready [get]
func Ready(c *gin.Context) {
	handlers.Ready(c)
}
