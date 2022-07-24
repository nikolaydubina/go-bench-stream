package gostreambench_test

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"
	"testing"

	gostreambench "github.com/nikolaydubina/go-bench-stream"
)

func ExampleIteratorSelector() {
	const sColorsOK = `
red
green
red
brown
yellow
brown
purple
red
`

	var p gostreambench.Iterator = gostreambench.NewIteratorSelectorFromReader(strings.NewReader(sColorsOK))

	p = gostreambench.IteratorSelector{It: p, Vals: map[string]bool{"red": true, "green": true, "brown": true}}
	p = gostreambench.IteratorSelector{It: p, Vals: map[string]bool{"red": true, "green": true}}

	b, err := io.ReadAll(gostreambench.NewReaderFromIteratorSelector(p))
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Print(string(b))
	// Output:
	// red
	// green
	// red
	// red
}

func benchIteratorSelector(b *testing.B, n int) {
	var s string

	for i := 0; i < n; i++ {
		s += dict[rand.Intn(len(dict))] + "\n"
	}

	var o []byte
	var err error

	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		var p gostreambench.Iterator = gostreambench.NewIteratorSelectorFromReader(strings.NewReader(s))

		p = gostreambench.IteratorSelector{It: p, Vals: map[string]bool{"red": true, "green": true, "brown": true}}
		p = gostreambench.IteratorSelector{It: p, Vals: map[string]bool{"red": true, "green": true}}

		o, err = io.ReadAll(gostreambench.NewReaderFromIteratorSelector(p))
		if err != nil {
			b.Error(err)
		}
		if len(o) == 0 {
			b.Error("empty output")
		}
	}
}

func BenchmarkIteratorSelector(b *testing.B) {
	for _, q := range []int{100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("n=%d", q), func(b *testing.B) {
			benchIteratorSelector(b, q)
		})
	}
}

func benchIteratorSelectorChain(b *testing.B, n int, l int) {
	var s string

	for i := 0; i < n; i++ {
		s += dict[rand.Intn(len(dict))] + "\n"
	}

	f := map[string]bool{"red": true, "green": true, "brown": true}

	var o []byte
	var err error

	b.ResetTimer()

	for n := 0; n < b.N; n++ {

		var p gostreambench.Iterator = gostreambench.NewIteratorSelectorFromReader(strings.NewReader(s))

		for j := 0; j < l; j++ {
			p = gostreambench.IteratorSelector{It: p, Vals: f}
		}

		o, err = io.ReadAll(gostreambench.NewReaderFromIteratorSelector(p))
		if err != nil {
			b.Error(err)
		}
		if len(o) == 0 {
			b.Error("empty output")
		}
	}
}

func BenchmarkIteratorSelector_Chain(b *testing.B) {
	for _, l := range []int{2, 4, 8, 16, 32, 64, 128, 256} {
		for _, n := range []int{100, 1000, 10000, 100000} {
			b.Run(fmt.Sprintf("l=%d_n=%d", l, n), func(b *testing.B) {
				benchIteratorSelectorChain(b, n, l)
			})
		}
	}
}
