package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Este metodo significa que el tipo T esta implementando la interface I
// Pero no necesitamos declarar explicitamente lo que hace
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
