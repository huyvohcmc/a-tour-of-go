package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{2, 4}
	p := &v
	p.X = 88
	fmt.Println(v)
}
