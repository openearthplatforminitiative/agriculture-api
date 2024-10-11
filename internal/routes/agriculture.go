package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/openearthplatforminitiative/agriculture-api/handlers"
	"github.com/openearthplatforminitiative/agriculture-api/models"
	"net/http"
)

func Agriculture(c *gin.Context) {
	response := handlers.Summary(c)
	var soilData models.SoilTypeJSON
	if err := json.Unmarshal(response, &soilData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse weather data"})
		return
	}
	c.JSON(http.StatusOK, soilData)
}
