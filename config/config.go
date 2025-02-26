package config

import (
	"os"
	"strconv"
)

type Config struct {
	Version        string
	ServerBindPort int
	ServerBindHost string
	ApiRootPath    string
	ApiDescription string
	ApiDomain      string
	ApiBaseUrl     string
}

func (c *Config) GetServerBindAddress() string {
	return c.ServerBindHost + ":" + strconv.Itoa(c.ServerBindPort)
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultValue
}

var AppSettings = &Config{}

func Setup() {
	AppSettings = &Config{
		Version:        getEnv("VERSION", "0.0.1"),
		ServerBindPort: getEnvInt("SERVER_BIND_PORT", 8080),
		ServerBindHost: getEnv("SERVER_BIND_HOST", "0.0.0.0"),
		ApiRootPath:    getEnv("API_ROOT_PATH", ""),
		ApiBaseUrl:     getEnv("API_BASE_URL", "https://api.openepi.io"),
	}
}
