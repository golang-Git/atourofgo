//defer

package main

import "fmt"

func main() {
	//deferは関数の終了まで処理を遅延させる
	defer fmt.Println("world")

	fmt.Println("hello")
}
