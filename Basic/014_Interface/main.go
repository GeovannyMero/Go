package main

import "fmt"

// Definicion de la interface
type geometry interface {
	area() float64
	perim() float64
}

// una structura
type rect struct {
	width, height float64
}

// Implementar los metodos de la interface a la structura
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	measure(r)

}
