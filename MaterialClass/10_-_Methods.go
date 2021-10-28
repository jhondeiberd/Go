package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) DoIt() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func foo(v Vertex) float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3.4, 4}
	fmt.Println(v.DoIt())
	fmt.Println(foo(v))
}
