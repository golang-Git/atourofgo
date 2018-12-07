/*
if文
ifは条件式を簡易的に書くことができ、その際の文法は
if [簡易文];[条件式] {}　となる
構文の中で定義した変数はその構文の中でのみ有効という制限を利用し
x, y := 3, 5
if n := x * y; n % 2 == 0 {
	fmt.Printf("n(%d) is even\n", n)
} else {
	fmt.Printf("n(%d) is odd\n", n)
}
等のように演算の結果を一時的に格納するために利用したり、
if _, err := doSomething(); err != nil {
	//関数doSomething()がエラーだった場合の処理
}
のように記述し、dosomething()で2番目の戻り値にエラーの有無を返すように定義しておくことで
エラーの処理をif文で補足することができる
*/
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
