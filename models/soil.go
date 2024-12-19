package models

type SoilType = SoilTypeInfo

// Below is the structure of the JSON from the weather API.
// It is only the sub-structure of the data.
// This way, we only get what we need when unmarshalling the JSON.

type SoilTypeJSON struct {
	// The soil type information at the queried location
	Properties SoilTypeInfo `json:"properties"`
}

type SoilTypeInfo struct {
	// The most probable soil type at the queried location
	MostProbableSoilType string `json:"most_probable_soil_type,omitempty"`
	Error                string `json:"error,omitempty"`
}
