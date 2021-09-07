package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.y*v.y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
