package env

import (
	"os"
	"strconv"
)

// TODO improve this code
func GetString(key, fallback string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	} else if val == "" {
		return fallback
	}
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	vaAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return vaAsInt
}
