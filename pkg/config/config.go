package config

import "os"

func Get(key string) string {
	return os.Getenv(key)
}

func GetWithFallback(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
