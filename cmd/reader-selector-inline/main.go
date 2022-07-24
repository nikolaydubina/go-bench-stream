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
		dict1Str string
		dict2Str string
	)
	flag.StringVar(&dict1Str, "dict1", "", "comma separated list of words, e.g. red,green,yellow,brown,purple")
	flag.StringVar(&dict2Str, "dict2", "", "comma separated list of words, e.g. red,green,yellow,brown,purple")
	flag.Parse()

	var dict1, dict2 map[string]bool
	dict1 = newdict(dict1Str)
	dict2 = newdict(dict2Str)

	var p io.Reader
	p = os.Stdin
	p = gostreambench.NewReaderSelector(p, dict1)
	p = gostreambench.NewReaderSelector(p, dict2)

	if _, err := io.Copy(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
}

func newdict(s string) map[string]bool {
	dict := make(map[string]bool, len(s))
	for _, q := range strings.Split(s, ",") {
		if q == "" {
			log.Fatalf("empty val in dict")
		}
		dict[q] = true
	}
	if len(dict) == 0 {
		log.Fatalf("empty dict")
	}
	return dict
}
