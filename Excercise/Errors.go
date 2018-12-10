package main

import (
	"fmt"
)

/*
Excercise:Errors.go
https://play.golang.org/p/--1_BMU1xN9

*/

type ErrNegativeSqrt float64

//Error()のeをキャストするのはeを単体で呼ぶとErrNegaticeSqrt型なのでErrorメソッドが再帰的に呼ばれてしまうため。
//⇒Error()メソッドを定義した型の場合は強制的にError()メソッドが呼ばれる(？)
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", uint(e))
}

//Sqrt is
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
