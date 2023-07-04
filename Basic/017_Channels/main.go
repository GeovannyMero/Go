// Los canales son tuberias que conectan go rutinas
// Se puede enviar valores en un canal desde una go rutina y recibirla desde otra go rutina.
// EL operador es el siguiente <-
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // Se envia sum a c (canal)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // Recive desde c

	fmt.Println(x, y, x+y)

}

// Para crear un canal se usa la siguiente sentencia make(chan val-type)
