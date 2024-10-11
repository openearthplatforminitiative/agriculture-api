package models

type SoilTypeJSON struct {
	// The feature type of the geojson-object
	Type interface{} `json:"type"`
	// The geometry of the queried location
	Geometry PointGeometry `json:"geometry"`
	// The soil type information at the queried location
	Properties SoilTypeInfo `json:"properties"`
}

type PointGeometry struct {
	// [longitude, latitude] decimal coordinates
	Coordinates []float32 `json:"coordinates"`
}

type SoilTypeInfo struct {
	// The most probable soil type at the queried location
	MostProbableSoilType string `json:"most_probable_soil_type"`
}
