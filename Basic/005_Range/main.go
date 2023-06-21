package main

import "fmt"

func main() {
	fmt.Println("Range")

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// i => indice
	// v => elemento
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// omitir los valores
	for i, _ := range pow {
		fmt.Println(i)
	}

	// Omitir los indices
	for _, v := range pow {
		fmt.Println(v)
	}

	// MAP ITERATIONS: Key and values
	fmt.Println("MAP")

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
