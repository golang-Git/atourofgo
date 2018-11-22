package main

import "fmt"

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}

func main() {
	returnFunc()()
	f := returnFunc()
	f()
	/*
		callFunction(func() {
			fmt.Println("Hello World!")
		})
	*/
	fmt.Printf("returnFunc=%T f=%T\n", returnFunc, f)
	callFunction(returnFunc())
}
