build:
	go build -o bin/gen-rand-dict ./cmd/gen-rand-dict
	go build -o bin/reader-selector ./cmd/reader-selector 
	go build -o bin/reader-selector-inline ./cmd/reader-selector-inline

clean-data:
	-rm -rf testdata
	mkdir testdata

data: clean-data
	./gen-rand-dict -n 1000 -dict red,green,yellow,brown,purple > testdata/colors-basic-1000.csv
	./gen-rand-dict -n 1000000 -dict red,green,yellow,brown,purple > testdata/colors-basic-1000000.csv
	./gen-rand-dict -n 100000000 -dict red,green,yellow,brown,purple > testdata/colors-basic-100000000.csv

validate:
	diff -q testdata/colors-basic-100000000.csv.out.reader-selector testdata/colors-basic-100000000.csv.out.grep
	diff -q testdata/colors-basic-100000000.csv.out.reader-selector-inline testdata/colors-basic-100000000.csv.out.grep

clean-profile:
	-rm -rf profiles
	mkdir profiles

bench: clean-profile
	go test -bench=. -benchtime=10s -benchmem .
