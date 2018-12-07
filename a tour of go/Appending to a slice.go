package main

import "./slice"

func main() {
	var s []int
	slice.Slice(s)

	s = append(s, 0)
	slice.Slice(s)

	s = append(s, 1)
	slice.Slice(s)

	s = append(s, 2, 3, 4)
	slice.Slice(s)

	x := s[:]
	slice.Slice(x)
}
