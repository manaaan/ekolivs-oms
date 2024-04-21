package env

import (
	"log"
	"os"
)

func Required(e string) string {
	value, ok := os.LookupEnv(e)
	if !ok {
		log.Fatalf("No value for env var %s", e)
	}

	return value
}
