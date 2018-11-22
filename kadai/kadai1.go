/*
🤔もんだい 1
for 文を使って 1 から 10 までの合計 (総和) を求めましょう。
1+2+3...+10 = 55
*/
package main

import "fmt"

func main() {
	// 😤
	var sum int = 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
