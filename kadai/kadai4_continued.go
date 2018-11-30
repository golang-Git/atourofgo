package main

import (
	"fmt"
)

/*
🤔もんだい 4: おいうえあ

あいうえお という文字列をパッケージ関数を利用せず、
逆順に(反転させて)出力してください。

*/

func main() {
	str := "あいうえお"

	fmt.Printf(Reverse(str))
}

func Reverse(str string) string {
	// 😤
	runes := []rune(str)
	length := len(runes)
	work := make([]rune, length)

	for i, j := 0, length; j > 0; i, j = i+1, j-1 {
		work[i] = runes[j-1]
	}
	return string(work)
}
