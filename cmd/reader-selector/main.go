package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	gostreambench "github.com/nikolaydubina/go-bench-stream"
)

func main() {
	var (
		dictStr string
	)
	flag.StringVar(&dictStr, "dict", "", "comma separated list of words, e.g. red,green,yellow,brown,purple")
	flag.Parse()

	dict := make(map[string]bool, len(dictStr))
	for _, q := range strings.Split(dictStr, ",") {
		if q == "" {
			log.Fatalf("empty val in dict")
		}
		dict[q] = true
	}
	if len(dict) == 0 {
		log.Fatalf("empty dict")
	}

	p := gostreambench.NewReaderSelector(os.Stdin, dict)
	if _, err := io.Copy(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
}
