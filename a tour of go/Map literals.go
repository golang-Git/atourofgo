package main

import "fmt"

type Vertex struct {
	lat, lng float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -122.08408},
	"Google":    {37.42202, -122.08 / 408},
}

func main() {
	fmt.Print(m)
}
