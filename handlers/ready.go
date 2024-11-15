package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/openearthplatforminitiative/agriculture-api/config"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
)

func Ready(c *gin.Context) {
	endpoints := []string{
		"/soil",
		"/weather",
		"/flood",
		"/deforestation",
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var failed bool

	for _, endpoint := range endpoints {
		wg.Add(1)
		go func(endpoint string) {
			defer wg.Done()
			if !checkHealth(getHealth(endpoint)) {
				mu.Lock()
				failed = true
				mu.Unlock()
			}
		}(endpoint)
	}

	wg.Wait()

	if failed {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Some endpoints are not ready"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "The service is ready"})
	}
}

func getHealth(endpoint string) []byte {
	base, err := url.Parse(config.AppSettings.ApiBaseUrl + endpoint + "/ready")
	if err != nil {
		log.Println("Failed to parse base URL:", err)
	}

	resp, err := http.Get(base.String())
	if err != nil {
		log.Println("Failed to call", endpoint, ":", err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close response body from", endpoint, ":", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Println("Received non-OK response from", endpoint, ":", resp.Status)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body from", endpoint, ":", err)
		return nil
	}

	return body
}

type HealthResponse struct {
	Hostname string `json:"hostname"`
	Status   string `json:"status"`
}

func checkHealth(body []byte) bool {
	var response HealthResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Failed to unmarshal JSON:", err)
	}

	return response.Status == "success"
}
