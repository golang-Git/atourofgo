//Slice literals

package main

import "fmt"

func main() {

	a := []int{2, 3, 5, 7, 11, 13}
	b := []bool{true, false, true, true, false, true}

	s := []struct {
		x int
		y bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s)
}
