package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// c := make(chan int, 1) // com buffer
	// c <- 42 // com buffer

	go func () {
		c <- 42
	}()

	fmt.Println(<-c)
}


// - Nível 10?! Êita! Parabéns!
// - Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
//     - Usando uma função anônima auto-executável
//     - Usando buffer