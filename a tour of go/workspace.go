package main

import (
	"fmt"
)

func main() {
	const (
		//Pi = 3.14159265358979323846264338327950288419716939937510582097494459
		F = 1.0000000000001
	)
	/*
		fmt.Println(float32(Pi))
		fmt.Println(float64(Pi))
		fmt.Println(Pi)

		a := float32(Pi) * 1000
		b := float64(Pi) * 1000
		c := Pi * 1000

		fmt.Printf("float32(Pi) * 1000 = %v\n", a)
		fmt.Printf("float64(Pi) * 1000 = %v\n", b)
		fmt.Printf("Pi * 1000 = %v\n", c)

	*/
	fmt.Println()
	fmt.Println(float64(F))
	fmt.Println(F)

	fmt.Println(float64(F) * 10000)
	fmt.Println(F * 10000)
}
