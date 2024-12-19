package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/openearthplatforminitiative/agriculture-api/config"
	"github.com/openearthplatforminitiative/agriculture-api/models"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Ready(c *gin.Context) {
	soilChan := make(chan models.HealthCheckResult)
	weatherChan := make(chan models.HealthCheckResult)
	floodChan := make(chan models.HealthCheckResult)
	deforestationChan := make(chan models.HealthCheckResult)

	go func() { soilChan <- getHealth("/soil") }()
	go func() { weatherChan <- getHealth("/weather") }()
	go func() { floodChan <- getHealth("/flood") }()
	go func() { deforestationChan <- getHealth("/deforestation") }()

	soilData := <-soilChan
	weatherData := <-weatherChan
	floodData := <-floodChan
	deforestationData := <-deforestationChan

	results := []models.HealthCheckResult{soilData, weatherData, floodData, deforestationData}

	c.JSON(http.StatusOK, models.HealthCheckResultSummary{Results: results})
}

func getHealth(endpoint string) models.HealthCheckResult {
	base, err := url.Parse(config.AppSettings.ApiBaseUrl + endpoint + "/ready")
	if err != nil {
		log.Println("Failed to parse base URL:", err)
		return models.HealthCheckResult{Endpoint: endpoint, Status: "failed", Error: err.Error()}
	}

	resp, err := http.Get(base.String())
	if err != nil {
		log.Println("Failed to call", endpoint, ":", err)
		return models.HealthCheckResult{Endpoint: endpoint, Status: "failed", Error: err.Error()}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close response body from", endpoint, ":", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Println("Received non-OK response from", endpoint, ":", resp.Status)
		return models.HealthCheckResult{Endpoint: endpoint, Status: "failed", Error: resp.Status}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body from", endpoint, ":", err)
		return models.HealthCheckResult{Endpoint: endpoint, Status: "failed", Error: err.Error()}
	}

	var response models.HealthResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Failed to unmarshal JSON:", err)
		return models.HealthCheckResult{Endpoint: endpoint, Status: "failed", Error: err.Error()}
	}

	if response.Status != "success" {
		return models.HealthCheckResult{Endpoint: endpoint, Status: "failed", Error: "Non-success status"}
	}

	return models.HealthCheckResult{Endpoint: endpoint, Status: "success"}
}
