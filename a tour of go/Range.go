package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 8, 16, 32, 64, 128}
	//rangeはforループなどで利用すると、スライスの場合、反復ごとに二つの変数を返す
	//1つは、添え字、もう1つはその添え字が指す場所の要素のコピー
	for i, v := range s {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
