package main

import "fmt"

func main() {
	var timeSone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
	}

	fmt.Println(timeSone["EST"])

}
