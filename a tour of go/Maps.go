package main

import "fmt"

type Vertex struct {
	lat, lng float64
}

//関数の外で変数を定義する場合、暗黙的な定義は利用できない
var m map[string]Vertex

func main() {

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Print(m["Bell Labs"])
}
