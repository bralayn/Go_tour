package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.55, -73.89},
	"Google":    {37.402, -122.87},
}

func main() {
	fmt.Println(m)
}
