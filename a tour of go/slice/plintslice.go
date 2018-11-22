package slice

import "fmt"

func Printslice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %d\n", s, len(x), cap(x), x)
}
