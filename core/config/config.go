package config

import (
	"main/core/utils"
	"strconv"
)

type Config struct {
	Port        int
	Timeout     int
	DatabaseURI string
}

func GetConfig() Config {
	return Config{
		Port:        parseEnvToInt("PORT", "5000"),
		Timeout:     parseEnvToInt("TIMEOUT", "30"),
		DatabaseURI: utils.GetEnv("DATABASE_URI", "http://localhost:8529"),
	}
}

func parseEnvToInt(envName, defaultValue string) int {
	num, err := strconv.Atoi(utils.GetEnv(envName, defaultValue))
	if err != nil {
		return 0
	}
	return num
}
