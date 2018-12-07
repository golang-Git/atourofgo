package main

import (
	"fmt"
	"strings"
)

func main() {
	s := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	s[0][0] = "X"
	s[2][2] = "O"
	s[1][2] = "X"
	s[1][0] = "O"
	s[0][2] = "X"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}
