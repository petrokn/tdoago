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

func InitLogging() {
	logFile, err := os.OpenFile("localization_log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)
}
