package models

type Summary struct {
	// The most probable soil type in the given coordinates.
	Soil SoilType `json:"soil"`

	// Current weather and rain forecast in the given coordinates.
	Weather Weather `json:"weather"`

	// Flood forecast in the given coordinates.
	Flood Flood `json:"flood"`

	// deforestation from 2001 to 2022 in the given coordinates.
	Deforestation Deforestation `json:"deforestation"`
}
