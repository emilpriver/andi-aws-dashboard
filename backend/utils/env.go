package utils

import (
	"os"
)

func EnvVariable(key string) string {
	return os.Getenv(key)
}
