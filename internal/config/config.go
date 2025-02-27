// Package config provides the configuration for the application.
package config

import (
	"os"
	"strconv"
)

// Config armazena todas as configurações da aplicação
type Config struct {
	DatabaseURL    string
	Port           int
	Auth0Domain    string
	Auth0Audience  string
	Env            string
	MigrationsPath string
}

// Load carrega as configurações das variáveis de ambiente
func Load() *Config {
	return &Config{
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/kidsdb?sslmode=disable"),
		Port:           getEnvInt("PORT", 8080),
		Auth0Domain:    getEnv("AUTH0_DOMAIN", ""),
		Auth0Audience:  getEnv("AUTH0_AUDIENCE", ""),
		Env:            getEnv("ENV", "development"),
		MigrationsPath: getEnv("MIGRATIONS_PATH", "./internal/migrations"),
	}
}

// Funções auxiliares para ler variáveis de ambiente
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	strValue := getEnv(key, "")
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return defaultValue
}
