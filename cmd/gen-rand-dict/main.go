package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
)

const doc string = `
Generate new line separated random sequence of words from dictionary.

Example:
$ gen-rand-dict -n 100 -dict red,green,yellow,brown,purple

Command options:
`

func main() {
	var (
		n       uint
		dictStr string
	)

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), doc)
		flag.PrintDefaults()
	}

	flag.UintVar(&n, "n", 10, "number of lines output")
	flag.StringVar(&dictStr, "dict", "", "comma separated list of words, e.g. red,green,yellow,brown,purple")
	flag.Parse()

	dict := strings.Split(dictStr, ",")
	if len(dict) == 0 {
		log.Fatalf("empty dict")
	}
	for _, q := range dict {
		if q == "" {
			log.Fatalf("empty val in dict")
		}
	}

	for i := uint(0); i < n; i++ {
		fmt.Println(dict[rand.Intn(len(dict))])
	}
}
