package env

import (
	"github.com/kynmh69/mormorare/pkg/logging"
	"os"
)

func FindEnv(key string) string {
	logger := logging.GetLogger()
	logger.Debug("Finding environment variable: ", key)
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	logger.Panic("Environment variable not found: ", key)
	return ""
}
