package main

import (
	"fmt"
	"log"
	"os"
	"tdoago/src/configuration"
	"time"

	"github.com/youpy/go-riff"
	"github.com/youpy/go-wav"
	"tdoago/src/common"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: <file.wav>\n")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	common.Check(err)

	config := configuration.Configure("config.yaml")

	common.InitLogging(&config.CoreConfiguration.LogFile)

	startTime := time.Now()

	reader := riff.Reader{file}

	dec := wav.NewReader(reader.RIFFReader)

	samples, err := dec.ReadSamples()
	common.Check(err)

	fmt.Println(samples)

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)

	log.Printf("Elapsed %.5f", elapsed.Seconds())
}
