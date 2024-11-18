package models

type Weather struct {
	AirTemperature      float32 `json:"air_temperature,omitempty"`
	CloudAreaFraction   float32 `json:"cloud_area_fraction,omitempty"`
	RelativeHumidity    float32 `json:"relative_humidity,omitempty"`
	WindFromDirection   float32 `json:"wind_from_direction,omitempty"`
	WindSpeed           float32 `json:"wind_speed,omitempty"`
	WindSpeedOfGust     float32 `json:"wind_speed_of_gust,omitempty"`
	PrecipitationAmount float32 `json:"precipitation_amount,omitempty"`
	Error               string  `json:"error,omitempty"`
}

// Below is the structure of the JSON from the weather API.
// It is only the sub-structure of the data.
// This way, we only get what we need when unmarshalling the JSON.

type METJSONForecast struct {
	Properties WeatherApiModelsMetWeatherTypesForecast `json:"properties"`
}

type WeatherApiModelsMetWeatherTypesForecast struct {
	Timeseries []ForecastTimeStep `json:"timeseries"`
}

type ForecastTimeStep struct {
	// Forecast for a specific time
	Data WeatherJSON `json:"data"`
}

type WeatherJSON struct {
	// Parameters which applies to this exact point in time
	Instant     Instant     `json:"instant"`
	Next12Hours Next12Hours `json:"next_12_hours"`
}

type Instant struct {
	Details ForecastTimeInstant `json:"details"`
}

type Next12Hours struct {
	Details ForecastTimePeriod `json:"details"`
}

type ForecastTimePeriod struct {
	PrecipitationAmount float32 `json:"precipitation_amount"`
}

type ForecastTimeInstant struct {
	AirTemperature    float32 `json:"air_temperature"`
	CloudAreaFraction float32 `json:"cloud_area_fraction"`
	RelativeHumidity  float32 `json:"relative_humidity"`
	WindFromDirection float32 `json:"wind_from_direction"`
	WindSpeed         float32 `json:"wind_speed"`
	WindSpeedOfGust   float32 `json:"wind_speed_of_gust"`
}
