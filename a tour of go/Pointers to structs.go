//Pointers to structs

package main

import "fmt"

type hoge struct {
	X int
	Y int
}

func main() {
	fuga := hoge{1, 2}
	x := &fuga
	x.X = 1e9
	fmt.Println(*x)
}
