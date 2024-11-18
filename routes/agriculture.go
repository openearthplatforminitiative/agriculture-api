package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/openearthplatforminitiative/agriculture-api/handlers"
)

// @BasePath /api/v1

// Agriculture godoc
// @Summary Get summary of agriculture data
// @Description Gets aggregated data from Deforestation, Flood, Weather and Soil APIs. See the [API documentation](https://developer.openepi.io) for more information.
// @Description There is an error field for each of the APIs. If a certain API does not respond, it will instead return with an error message in the field that belong to that API.
// @Tags summary
// @Produce json
// @Param   lat     query    string     true        "Latitude"
// @Param   lon     query    string     true        "Longitude"
// @Success 200 {object} models.Summary
// @Router /summary [get]
func Agriculture(c *gin.Context) {
	handlers.Summary(c)
}
