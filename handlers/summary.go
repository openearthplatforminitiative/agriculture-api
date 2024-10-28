package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/openearthplatforminitiative/agriculture-api/models"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Summary(c *gin.Context) (summary models.Summary) {
	params := c.Request.URL.Query()

	// Check if we have the required parameters
	if params.Get("lat") == "" || params.Get("lon") == "" {
		log.Println("Missing required parameters")
		return
	}

	soilChan := make(chan []byte)
	weatherChan := make(chan []byte)
	floodChan := make(chan []byte)
	deforestationChan := make(chan []byte)

	go func() { soilChan <- getData(params, "soil/type") }()
	go func() { weatherChan <- getData(params, "weather/locationforecast") }()
	go func() { floodChan <- getData(params, "flood/summary") }()
	go func() { deforestationChan <- getData(params, "deforestation/basin") }()

	soilData := <-soilChan
	weatherData := <-weatherChan
	floodData := <-floodChan
	deforestationData := <-deforestationChan

	return createSummary(soilData, weatherData, floodData, deforestationData)
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

func createSummary(soilData, weatherData, floodData, deforestationData []byte) (summary models.Summary) {
	var soil models.SoilTypeJSON
	var weather models.METJSONForecast
	var flood models.SummaryResponseModel
	var deforestation models.DeforestationBasinGeoJSON

	if err := json.Unmarshal(soilData, &soil); err != nil {
		log.Println("Failed to unmarshal soil data:", err)
		return
	}
	if err := json.Unmarshal(weatherData, &weather); err != nil {
		log.Println("Failed to unmarshal weather data:", err)
		return
	}
	if err := json.Unmarshal(floodData, &flood); err != nil {
		log.Println("Failed to unmarshal flood data:", err)
		return
	}
	if err := json.Unmarshal(deforestationData, &deforestation); err != nil {
		log.Println("Failed to unmarshal deforestation data:", err)
		return
	}

	var firstWeather models.ForecastTimeInstant
	var precipitationAmount float32

	if len(weather.Properties.Timeseries) > 0 {
		firstWeather = weather.Properties.Timeseries[0].Data.Instant.Details
		precipitationAmount = weather.Properties.Timeseries[0].Data.Next12Hours.Details.PrecipitationAmount
	}

	var firstFlood models.SummaryFeature
	if len(flood.QueriedLocation.Features) > 0 {
		firstFlood = flood.QueriedLocation.Features[0]
	}

	var firstDeforestation models.DeforestationBasinFeature
	if len(deforestation.Features) > 0 {
		firstDeforestation = deforestation.Features[0]
	}

	body := models.Summary{
		MostProbableSoilType: soil.Properties.MostProbableSoilType,
		Weather: models.Weather{
			AirTemperature:      firstWeather.AirTemperature,
			CloudAreaFraction:   firstWeather.CloudAreaFraction,
			RelativeHumidity:    firstWeather.RelativeHumidity,
			WindFromDirection:   firstWeather.WindFromDirection,
			WindSpeed:           firstWeather.WindSpeed,
			WindSpeedOfGust:     firstWeather.WindSpeedOfGust,
			PrecipitationAmount: precipitationAmount,
		},
		Flood: firstFlood.Properties,
		Deforestation: models.Deforestation{
			DaterangeTotTreeloss: firstDeforestation.Properties.DaterangeTotTreeloss,
		},
	}

	return body
}
