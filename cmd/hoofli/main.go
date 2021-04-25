package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dnnrly/hoofli"
)

func main() {
	var harFile string
	flag.StringVar(&harFile, "har", "", "the location of a HAR file to parse")

	flag.Parse()

	if harFile == "" {
		fmt.Fprintln(os.Stderr, "must specify HAR input")
		os.Exit(1)
	}

	hf, _ := os.Open(harFile)
	defer hf.Close()

	har, _ := hoofli.NewHar(hf)
	_ = har.Draw(os.Stdout)
}
