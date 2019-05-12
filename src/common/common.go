package common

import (
	"log"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func InitLogging(filename *string) {
	logFile, err := os.OpenFile(*filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	log.SetOutput(logFile)
}
