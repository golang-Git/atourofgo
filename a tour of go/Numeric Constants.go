//Numeric Constants

package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func hoge(x int) int {
	x = x*10 + 1
	return x
}

func fuga(x float64) float64 {
	x = x * 0.1
	return x
}

func main() {
	fmt.Println(Small)
	fmt.Println(hoge(Small))
	fmt.Println(fuga(Small))
	fmt.Println(fuga(Big))
}
