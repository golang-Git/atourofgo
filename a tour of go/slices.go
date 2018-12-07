//slices

package main

import "fmt"

func main() {
	a := [10]int{2, 3, 5, 7, 11, 13}
	prime := a[1:4] //var s []int = primes[1:4]
	fmt.Println(prime)
}
