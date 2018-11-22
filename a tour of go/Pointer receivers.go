package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
レシーバにポインターを利用することでレシーバの変数の値を変更することができる
関数Scale()のレシーバをポインタ型ではなくVertex型に変更するとScaleの中で代入した値はmainの中では適用されない
*/
func (v *Vertex) Scale(i float64) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
