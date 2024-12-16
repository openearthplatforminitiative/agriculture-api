package models

type HealthCheckResultSummary struct {
	Results []HealthCheckResult `json:"results"`
}

type HealthCheckResult struct {
	Endpoint string `json:"endpoint"`
	Status   string `json:"status"`
	Error    string `json:"error,omitempty"`
}

type HealthResponse struct {
	Hostname string `json:"hostname"`
	Status   string `json:"status"`
}
