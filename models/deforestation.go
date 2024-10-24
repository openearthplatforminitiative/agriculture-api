package models

type Deforestation struct {
	DaterangeTotTreeloss float32 `json:"daterange_tot_treeloss"`
}

// Below is the structure of the JSON from the weather API.
// It is only the sub-structure of the data.
// This way, we only get what we need when unmarshalling the JSON.

type DeforestationBasinGeoJSON struct {
	Features []DeforestationBasinFeature `json:"features"`
}

type DeforestationBasinFeature struct {
	// Unique basin polygon identifier.
	Properties BasinProperties `json:"properties"`
}

type BasinProperties struct {
	// Total tree cover loss, in square kilometers, within the basin polygon over the time period from start_year to end_year (inclusive)
	DaterangeTotTreeloss float32 `json:"daterange_tot_treeloss"`
}
