package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

/*
Stringer型（interface）
文字列を返すメソッドString
定義することで暗黙的に読みやすい形式へ適切に変換してくれる
*/
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	//fmtで引数に使用されている型のStringメソッドが定義されている場合、文字列の出力に利用してくれる
	fmt.Println(a, z)
}
