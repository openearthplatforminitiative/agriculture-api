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

func Ready(c *gin.Context) {
	handlers.Ready(c)
}
