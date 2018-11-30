package main

import (
	"fmt"
)

/*
🤔もんだい 4: EDCBA

ABCDE という文字列をパッケージ関数を利用せず、
逆順に(反転させて)出力してください。

*/

func main() {
	str := "ABCDE"

	fmt.Printf(Reverse(str))
}

func Reverse(str string) (s string) {
	// 😤
	var arr [5]string
	for key, val := range str {
		arr[key] = string(val)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		s += arr[i]
	}

	return
}
