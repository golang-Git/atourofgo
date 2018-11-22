package main

import "fmt"

func main() {
	s := make([]int, 10)
	for i := range s {
		s[i] = 1 << uint(i) //shift演算 01 << 1 = 10 = 2
		//       　 01 << 2 = 100 = 4 …
	}

	for _, value := range s {
		fmt.Printf("%d\n", value)
	}
}
