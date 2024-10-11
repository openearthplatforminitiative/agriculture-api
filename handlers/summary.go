package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Summary(c *gin.Context) (body []byte) {
	params := c.Request.URL.Query()
	soilData := getData(params, "soil/type")
	weatherData := getData(params, "weather/locationforecast")
	floodData := getData(params, "flood/summary")
	deforestationData := getData(params, "deforestation/basin")

	// TODO: combine data into something useful. Create Summary.go model. We want a short summary of the conditions in the given coordinates

	return soilData
}

func getData(params url.Values, endpoint string) (body []byte) {
	resp, err := http.Get("https://api.openepi.io/" + endpoint + "?lat=" + params.Get("lat") + "&lon=" + params.Get("lon"))
	if err != nil {
		log.Println("Failed to fetch data")
		return
	}

	// Close the response body on function exit
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close response body")
			return
		}
	}(resp.Body)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body")
		return
	}
	return body
}
