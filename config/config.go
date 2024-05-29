package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type FileConfig struct {
	FilePath string `envconfig:"FILE_PATH"`
}

func MustParseConfig(envPath string, c interface{}) {
	err := godotenv.Load(envPath)
	if err != nil {
		log.WithError(err).Warn("failed to load .env file")
	}

	if err := envconfig.Process("", c); err != nil {
		log.WithError(err).Fatal("could not load config from ENV")
	}
}

func GetEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}
