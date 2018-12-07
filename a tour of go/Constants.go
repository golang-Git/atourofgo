//Constants

package main

import "fmt"

const Pi = 3.14 //定数(constant) 文字、文字列、boolean、数値で利用できる、定数はその状況によって必要な型を取ることができる

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Rule = true
	fmt.Println("Go rules?", Rule)
}
