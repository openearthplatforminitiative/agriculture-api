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

func Summary(c *gin.Context) {
	// Check if we have the required parameters
	if c.Query("lat") == "" || c.Query("lon") == "" {
		log.Println("Missing required parameters")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters: lat and lon"})
		return
	}

	soilChan := make(chan []byte)
	weatherChan := make(chan []byte)
	floodChan := make(chan []byte)
	deforestationChan := make(chan []byte)

	go func() { soilChan <- getData(c, "/soil/type") }()
	go func() { weatherChan <- getData(c, "/weather/locationforecast") }()
	go func() { floodChan <- getData(c, "/flood/summary") }()
	go func() { deforestationChan <- getData(c, "/deforestation/basin") }()

	soilData := <-soilChan
	weatherData := <-weatherChan
	floodData := <-floodChan
	deforestationData := <-deforestationChan

	summary, err := createSummary(soilData, weatherData, floodData, deforestationData)
	if err != nil {
		log.Println("Failed to create summary:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create summary"})
		return
	}

	c.JSON(http.StatusOK, summary)
}

func getData(c *gin.Context, endpoint string) []byte {
	base, err := url.Parse(config.AppSettings.ApiBaseUrl)
	if err != nil {
		log.Println("Failed to parse base URL:", err)
		return nil
	}

	base.Path += endpoint

	p := url.Values{}
	p.Set("lat", c.Query("lat"))
	p.Set("lon", c.Query("lon"))
	base.RawQuery = p.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		log.Println("Failed to fetch data from", endpoint, ":", err)
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

func createSummary(soilData, weatherData, floodData, deforestationData []byte) (models.Summary, error) {
	var soil models.SoilTypeJSON
	var weather models.METJSONForecast
	var flood models.SummaryResponseModel
	var deforestation models.DeforestationBasinGeoJSON

	if err := json.Unmarshal(soilData, &soil); err != nil {
		return models.Summary{}, err
	}
	if err := json.Unmarshal(weatherData, &weather); err != nil {
		return models.Summary{}, err
	}
	if err := json.Unmarshal(floodData, &flood); err != nil {
		return models.Summary{}, err
	}
	if err := json.Unmarshal(deforestationData, &deforestation); err != nil {
		return models.Summary{}, err
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

	return body, nil
}
