package main

import "fmt"

func main() {
	names := [4]string{
		"Jhon",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	//length y capacidad
	s := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(s)

	//Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	//Extend its length
	s = s[:4]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
