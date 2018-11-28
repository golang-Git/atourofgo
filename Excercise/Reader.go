package main

import (
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (i int, e error) {
	r = strings.NewReader("A")
	i, e = r.Read(b)
	return i, e
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	var read MyReader
	reader.Validate(read.Read)
}
