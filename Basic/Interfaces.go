package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := vertec{3, 4}

	a = f  //a MyFloat implementa Abser
	a = &v // a *Vertex implementa Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser
	//a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(64)
	}
	return float64(f)
}

type vertec struct {
	X, Y float64
}

func (v *vertec) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
