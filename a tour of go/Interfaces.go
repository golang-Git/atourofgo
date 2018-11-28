package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := myFloat(-math.Sqrt2)
	//v := Vertex{3, 4}

	a = f
	//a = &v

	//a = v

	fmt.Println(a.Abs())
}

type myFloat float64

type Vertex struct {
	X, Y float64
}

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
