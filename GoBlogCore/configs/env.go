package configs

import (
	"log"
	"os"
)

var (
	MongoUri = GetEnv("MONGO_URI")
	DbName   = GetEnv("DB_NAME")
	HttpPort = GetEnv("HTTP_PORT")
)

func GetEnv(name string) string {
	err := os.Getenv(name)
	if len(err) <= 0 {
		log.Fatal(name + " environment variable is missing!")
	}

	return os.Getenv(name)
}
