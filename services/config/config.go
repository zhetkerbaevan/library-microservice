package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ConnectionString string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load() //Load env variables
	return Config{
		ConnectionString: getEnv("CONNECTION_STRING", "mongodb://admin:1234@localhost:27017/library_db?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false"),
	}
}

func getEnv(key, fallback string) string {
	//Look for env variables by key
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback //Default
}
