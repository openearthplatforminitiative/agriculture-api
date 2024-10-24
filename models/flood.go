package models

type Flood = SummaryProperties

// Below is the structure of the JSON from the weather API.
// It is only the sub-structure of the data.
// This way, we only get what we need when unmarshalling the JSON.

// PeakTimingEnum the model 'PeakTimingEnum'
type PeakTimingEnum string

// List of PeakTimingEnum
const (
	BB PeakTimingEnum = "BB"
	GC PeakTimingEnum = "GC"
	GB PeakTimingEnum = "GB"
)

// IntensityEnum the model 'IntensityEnum'
type IntensityEnum string

// List of IntensityEnum
const (
	P IntensityEnum = "P"
	R IntensityEnum = "R"
	Y IntensityEnum = "Y"
	G IntensityEnum = "G"
)

type SummaryResponseModel struct {
	// A feature collection representing the queried location's summary forecast data.
	QueriedLocation SummaryFeatureCollection `json:"queried_location"`
}

type SummaryFeatureCollection struct {
	// A collection of summary forecasts, each containing specific forecast information for a queried location.
	Features []SummaryFeature `json:"features"`
}

type SummaryFeature struct {
	// Specific properties of the summary forecast, including various attributes like tendency, peak step, and intensity.
	Properties SummaryProperties `json:"properties"`
}

type SummaryProperties struct {
	// The date the summary forecast was issued on. The GloFAS hydrological model is run every day at 00:00 UTC.
	IssuedOn string `json:"issued_on"`
	// The step number at which the peak occurs, ranging from 1 to 30.
	PeakStep int32 `json:"peak_step"`
	// The date on which the flood peak is forecasted to occur, assuming UTC timezone.
	PeakDay string `json:"peak_day"`
	// The timing of the flood peak indicated by border and grayed colors. BB: Black border, peak forecasted within days 1-3. GC: Grayed color, peak forecasted after day 10 with <30% probability of exceeding the 2-year return period threshold in first 10 days. GB: Gray border, floods of some severity in first 10 days and peak after day 3.
	PeakTiming PeakTimingEnum `json:"peak_timing"`
	// The flood intensity (indicated by color) relating to maximum return period threshold exceedance probabilities over the forecast horizon. P: Purple, maximum 20-year exceedance probability >=30%; R: Red, maximum for 20-year <30% and 5-year >=30%; Y: Yellow, maximum for 5-year <30% and 2-year >=30%; G: Gray, no flood signal (2-year <30%).
	Intensity IntensityEnum `json:"intensity"`
}
