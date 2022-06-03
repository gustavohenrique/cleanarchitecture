package fileutils

import (
	"log"
	"os"
)

func logInfo(message ...interface{}) {
	debug := os.Getenv("DEBUG")
	if debug == "true" {
		log.Println("[INFO]", message)
	}
}
