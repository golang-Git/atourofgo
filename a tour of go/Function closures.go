package main

import (
	"fmt"
	_ "math" //利用していないパッケージについても_を先頭につけることでimportすることができる
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hoge, fuga := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(hoge(i), fuga(i*-2))
	}
}
