package main

import (
	"fmt"
)

/*
🤔もんだい 2: FizzBuzz
- ルール -
1~100 までの数字で、
3 で割り切れれば、「Fizz」を表示する
5 で割り切れれば、「Buzz」を表示する
3 と 5 で割り切れれば、「FizzBuzz」を表示する
上記以外の場合は、そのままの数字を表示する

算術演算子はコレを見ると良いよ！↓
http://golang.jp/go_spec#Arithmetic_operators
*/

func main() {
	// 😤
	for i := 1; i <= 100; i++ {
		fmt.Print(i)
		if (i % 3) == 0 {
			fmt.Print("Fizz")
		}
		if (i % 5) == 0 {
			fmt.Print("Buzz")
		}
		fmt.Println("")
	}
}
