//defer

package main

import "fmt"

func main() {
	//deferは関数の終了まで処理を遅延させる
	defer fmt.Println("world")

	fmt.Println("hello")

	//deferで複数の処理を登録したい場合は、無名関数を利用することができる
	defer func() {
		fmt.Println("A")
		fmt.Println("B")
		fmt.Println("C")
	}() //deferで呼び出せるのは関数呼び出し形式のものに限定されるため、関数として記述する場合は、最後に()をつける
}
