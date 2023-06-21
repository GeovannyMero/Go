package main

import "fmt"

func main() {
	x := 4
	if x < 4 {
		fmt.Println(x)
	} else {
		fmt.Println("OK")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
