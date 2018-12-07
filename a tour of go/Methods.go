package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	V, X float64
}

/*
methodを利用する場合はfuncと関数名の間にレシーバの型を記述する
定義したメソッドはレシーバ.メソッド()という形で呼び出すことができる

*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.V*v.V + v.X*v.X)
}

func main() {
	v := Vertex{4, 3}
	fmt.Println(v.Abs())
}
