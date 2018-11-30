package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err == nil {
		for i, v := range p {
			switch {
			case v >= 'a' && v <= 'z':
				p[i] = (v-'a'+13)%26 + 'a'
			case v >= 'A' && v <= 'Z':
				p[i] = (v-'A'+13)%26 + 'A'
			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
