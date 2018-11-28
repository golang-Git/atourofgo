package main

import (
	"strings"
	//"fmt"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	array := strings.Fields(s)
	for i := 0; i < len(array); i++ {
		if m[array[i]] == 0 {
			m[array[i]] = 1
		} else {
			m[array[i]] = (m[array[i]] + 1)
		}
	}
	//fmt.Print(m)
	return m
}

func main() {
	wc.Test(WordCount)
}
