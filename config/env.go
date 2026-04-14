package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var loadEnvOnce sync.Once

func loadEnv() {
	loadEnvOnce.Do(func() {
		// Ignore missing .env so deployed environments can rely on real env vars.
		_ = godotenv.Load()
	})
}

func getEnv(keys ...string) string {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}

	return ""
}
