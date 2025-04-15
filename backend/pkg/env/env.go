package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Required(e string) string {
	value, ok := os.LookupEnv(e)
	if !ok {
		log.Fatalf("No value for env var %s", e)
	}

	return value
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		// If injected on runtime, this is not an issue
		log.Println("Unable to load .env file")
	}
}
