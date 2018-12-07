//If and else

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

//出力は以下になるが、
//27 >= 20
//9 20
//実際のそれぞれの出力はpow(3, 3, 20),
//27 >= 20
//20
//とpow(3, 2, 10),
//9
//順番が混在しているのは関数はreturnし値が、mainのfmtの中に格納され、すべての関数の処理が終わるのを待つが関数中の出力はその場で返るため
