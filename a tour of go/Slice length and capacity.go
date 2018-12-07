package main

import "./slice"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	slice.Slice(s)
	s = s[:2]
	slice.Slice(s)
	s = s[:0]
	slice.Slice(s)
	s = s[:4]
	slice.Slice(s)
}
