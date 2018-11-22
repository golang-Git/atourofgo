package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var s, next, work int

	return func() int {
		work = s
		s = s + next
		next = work
		if s == 0 {
			next += 1
		}
		return s
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
