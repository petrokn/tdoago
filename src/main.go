package main

import (
	"fmt"
	"os"
	"tdoago/src/configuration"

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

	configuration.Configure("config.yaml")

	reader := riff.Reader{file}

	dec := wav.NewReader(reader.RIFFReader)

	samples, err := dec.ReadSamples()
	common.Check(err)

	fmt.Println(samples)
}
