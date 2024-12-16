// Package config menyediakan konfigurasi aplikasi
package config

import (
	"os"
	"strconv"
)

// Config menyimpan konfigurasi aplikasi
type Config struct {
	Port             string
	MaxSearchResults int
	ItemsPerPage     int
	DataPath         string
}

// LoadConfig memuat konfigurasi dari environment variables
func LoadConfig() *Config {
	return &Config{
		Port:             getEnv("PORT", "8080"),
		MaxSearchResults: getEnvAsInt("MAX_SEARCH_RESULTS", 100),
		ItemsPerPage:     getEnvAsInt("ITEMS_PER_PAGE", 10),
		DataPath:         getEnv("DATA_PATH", "articles.json"),
	}
}

// getEnv mengambil nilai environment variable dengan default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt mengambil nilai environment variable sebagai integer
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if v, err := strconv.Atoi(value); err == nil {
			return v
		}
	}
	return defaultValue
}

