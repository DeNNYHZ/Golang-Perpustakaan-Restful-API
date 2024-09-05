package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
	"time"
)

// LoadConfig loads environment variables from a .env file
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

// Config retrieves a configuration value based on the provided key.
func Config(key string) string {
	return os.Getenv(key)
}

// ServerTimeOut returns the server configuration including timeouts.
func ServerTimeOut() fiber.Config {
	return fiber.Config{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		// Add other configuration options if needed
	}
}
