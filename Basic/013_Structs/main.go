package main

import "fmt"

// A Struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Vertex{1, 2})

	fmt.Println(person{"Bob", 20})
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
