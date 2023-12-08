package config

import (
	"os"
	"strconv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	Db DbConfig
}

func GetConfig() *Config {
	return &Config{
		Db: DbConfig{
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_USER_PASSWORD", ""),
			Name:     getEnv("DB_NAME", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := getEnv(name, "")
	if value, err := strconv.Atoi(valStr); err == nil {
		return value
	}

	return defaultVal
}
