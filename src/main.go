package src

import (
	"fmt"
	"os"
	
	"github.com/youpy/go-wav"
	"github.com/youpy/go-riff"
	"tdoago/src/common"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: <file.wav>\n")
		os.Exit(1)
	}
	
	file, err := os.Open(os.Args[1])
	common.Check(err)

	reader := riff.Reader{file}

	dec := wav.NewReader(reader.RIFFReader)

	samples, err := dec.ReadSamples()
	common.Check(err)

	fmt.Println(samples)
}
