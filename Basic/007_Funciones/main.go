package main

import "fmt"

func main() {
	res := Plus(5, 8)
	fmt.Println(res)
}

func Plus(a int, b int) int {
	return a + b
}
