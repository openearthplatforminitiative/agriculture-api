package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"github.com/openearthplatforminitiative/agriculture-api/cmd/docs"
	"github.com/openearthplatforminitiative/agriculture-api/config"
	"github.com/openearthplatforminitiative/agriculture-api/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func init() {
	config.Setup()
	gin.SetMode(gin.ReleaseMode)
}

// @title           Agriculture API
// @version         1.0.0
// @description     This API is used to get aggregated data from Deforestation, Flood, Weather and Soil APIs.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()
	routes.InitRoutes(router)

	// Swagger documentation will not work without this even though it is defined over.
	docs.SwaggerInfo.BasePath = "/"

	// Serve static files for Swagger and Redoc documentation
	router.Static("/cmd/docs", "./cmd/docs")

	//Swagger endpoint
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Redoc endpoint
	router.GET("/redoc", func(c *gin.Context) {
		opts := middleware.RedocOpts{
			SpecURL: "/cmd/docs/swagger.json",
			Path:    "/redoc",
		}
		sh := middleware.Redoc(opts, nil)
		sh.ServeHTTP(c.Writer, c.Request)
	})

	log.Println("Starting server on", config.AppSettings.GetServerBindAddress())
	err := router.Run(config.AppSettings.GetServerBindAddress())
	if err != nil {
		log.Println("Failed to start server")
		return
	}
}
