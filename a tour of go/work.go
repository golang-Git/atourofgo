package main

import "fmt"

type mypint int

type myint int

func (i *mypint) hoge() int {
	return 1
}

func (j myint) fuga() int {
	return 2
}

func main() {
	a := mypint(2)
	j := myint(1)
	fmt.Println(a.hoge())
	fmt.Println(j.fuga())
}
