//Stacking defers

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		//deferに渡した関数が複数あるとLIFOで処理される
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
