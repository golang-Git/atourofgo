package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	//interfaceの中の具体的な値がnilの場合、メソッドはnilをレシーバとして呼び出す
	var i I = &T{"hello"}
	i.M()

	var t *T
	i = t
	i.M()

	i = F(math.Pi)
	i.M()
}
