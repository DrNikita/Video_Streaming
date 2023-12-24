package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type Config struct {
	AppPort       string   `map-structure:"APP_PORT"`
	FfmpegCommand string   `map-structure:"FFMPEG_ARGS"`
	AllowOrigins  []string `map-structure:"ALLOW_ORIGINS"`
}

var conf Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warning("Couldn't load env variables to connect to DB...")
	}

	conf = Config{
		AppPort:       getEnv("APP_PORT", "8080"),
		FfmpegCommand: getEnv("FFMPEG_ARGS", ""),
		AllowOrigins:  getEnvAsSlice("ALLOW_ORIGINS", []string{}, ","),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}

func GetConfigurationInstance() Config {
	return conf
}
