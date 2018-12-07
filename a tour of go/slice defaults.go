//silce defaults
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11}
	s = s[1:4]
	fmt.Println(s) //sの1番目から4番目の前までを取り出す({3,5,7})
	s = s[:2]
	fmt.Println(s) //s1-4のデータの2番目の前までを取り出す({3,5})
	s = s[1:]
	fmt.Println(s) //s1-2の1番目を取り出す{5}
}
