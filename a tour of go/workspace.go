package main

import "fmt"

//interface型の変数に対する型アサーション
func anything(a interface{}) {

	if a == nil {
		fmt.Println("a is nil")
	} else if i, isInt := a.(int); isInt {
		fmt.Printf("a is integer : %d\n", i)
	} else if s, isString := a.(string); isString {
		fmt.Println(s)
	} else {
		fmt.Println("unsupported type!")
	}
}

func main() {
	anything(1)
}
