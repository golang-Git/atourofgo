//If with a short statement

package main

import (
	"fmt"
	"math"
)

func pow(x, y, lim float64) float64 {
	if z := math.Pow(x, y); z < lim {
		return z
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}
