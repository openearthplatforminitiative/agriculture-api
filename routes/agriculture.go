package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/openearthplatforminitiative/agriculture-api/handlers"
	"net/http"
)

func Agriculture(c *gin.Context) {
	response := handlers.Summary(c)
	c.JSON(http.StatusOK, response)
}
