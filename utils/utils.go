package utils

import (
	"os"
)

func GetEnvOrDefault(env, str string) string {
	value := os.Getenv(env)
	if value == "" {
		value = str
	}
	return value
}
