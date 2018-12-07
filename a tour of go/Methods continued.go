package main

import (
	"fmt"
	"math"
)

type myFloat float64

/*
上記のように変数に対してエイリアスを行い、その型にメソッドを定義することもできる
*/
func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func main() {
	f := myFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
