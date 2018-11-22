//Structs

package main

import "fmt"

type V struct { //struct…構造体。複数のデータを一つにまとめることができる。
	X int
	Y int
}

func main() {
	v := V{1, 4}
	v.X = 3
	fmt.Println(v.X)
}
