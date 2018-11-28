package main

import "fmt"

//switch文を利用した型アサーションの方法

func main() {
	var v interface{} = 3

	switch s := v.(type) {
	case string, uint:
		//上記のようにcase文の中には複数の型を記述することができるが、その場合はcaseの中で計算などはできない
		//型が明確に定まる場合は、変数sを「その型そのものの変数」として利用できるが、複数の型を列挙したパターンでは変数vの型が1つに定まらない
	case int:
		fmt.Println("int", s)
	}
}
