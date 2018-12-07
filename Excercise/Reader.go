package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (i int, e error) {
	for i, e = 0, nil; n < len(b); i++ {
		b[i] = 'A'
	}
	return
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	var read MyReader
	reader.Validate(read.Read)
}
