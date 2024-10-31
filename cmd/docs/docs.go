// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/summary": {
            "get": {
                "description": "Gets aggregated data from Deforestation, Flood, Weather and Soil APIs. See the [API documentation](https://developer.openepi.io) for more information.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "summary"
                ],
                "summary": "Get summary of agriculture data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Latitude",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Longitude",
                        "name": "lon",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Summary"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Deforestation": {
            "type": "object",
            "properties": {
                "daterange_tot_treeloss": {
                    "type": "number"
                }
            }
        },
        "models.Flood": {
            "type": "object",
            "properties": {
                "intensity": {
                    "description": "The flood intensity (indicated by color) relating to maximum return period threshold exceedance probabilities over the forecast horizon. P: Purple, maximum 20-year exceedance probability \u003e=30%; R: Red, maximum for 20-year \u003c30% and 5-year \u003e=30%; Y: Yellow, maximum for 5-year \u003c30% and 2-year \u003e=30%; G: Gray, no flood signal (2-year \u003c30%).",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.IntensityEnum"
                        }
                    ]
                },
                "issued_on": {
                    "description": "The date the summary forecast was issued on. The GloFAS hydrological model is run every day at 00:00 UTC.",
                    "type": "string"
                },
                "peak_day": {
                    "description": "The date on which the flood peak is forecasted to occur, assuming UTC timezone.",
                    "type": "string"
                },
                "peak_step": {
                    "description": "The step number at which the peak occurs, ranging from 1 to 30.",
                    "type": "integer"
                },
                "peak_timing": {
                    "description": "The timing of the flood peak indicated by border and grayed colors. BB: Black border, peak forecasted within days 1-3. GC: Grayed color, peak forecasted after day 10 with \u003c30% probability of exceeding the 2-year return period threshold in first 10 days. GB: Gray border, floods of some severity in first 10 days and peak after day 3.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.PeakTimingEnum"
                        }
                    ]
                }
            }
        },
        "models.IntensityEnum": {
            "type": "string",
            "enum": [
                "P",
                "R",
                "Y",
                "G"
            ],
            "x-enum-varnames": [
                "P",
                "R",
                "Y",
                "G"
            ]
        },
        "models.PeakTimingEnum": {
            "type": "string",
            "enum": [
                "BB",
                "GC",
                "GB"
            ],
            "x-enum-varnames": [
                "BB",
                "GC",
                "GB"
            ]
        },
        "models.Summary": {
            "type": "object",
            "properties": {
                "deforestation": {
                    "description": "deforestation from 2001 to 2022 in the given coordinates.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Deforestation"
                        }
                    ]
                },
                "flood": {
                    "description": "Flood forecast in the given coordinates.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Flood"
                        }
                    ]
                },
                "most_probable_soil_type": {
                    "description": "The most probable soil type in the given coordinates.",
                    "type": "string"
                },
                "weather": {
                    "description": "Current weather and rain forecast in the given coordinates.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Weather"
                        }
                    ]
                }
            }
        },
        "models.Weather": {
            "type": "object",
            "properties": {
                "air_temperature": {
                    "type": "number"
                },
                "cloud_area_fraction": {
                    "type": "number"
                },
                "precipitation_amount": {
                    "type": "number"
                },
                "relative_humidity": {
                    "type": "number"
                },
                "wind_from_direction": {
                    "type": "number"
                },
                "wind_speed": {
                    "type": "number"
                },
                "wind_speed_of_gust": {
                    "type": "number"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Agriculture API",
	Description:      "This API is used to get aggregated data from Deforestation, Flood, Weather and Soil APIs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
