//Arrays

package main

import "fmt"

func main() {

	var a [2]string
	a[1] = "Hello"
	a[2] = "World"
	fmt.Println(a[1], a[2])
	fmt.Println(a)

	prime := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(prime)

}
