package env

import "os"

// Get env varilable by key or return provided fallback
func Get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
