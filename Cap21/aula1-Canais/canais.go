package main

import "fmt"

func main() {

	canal := make(chan int, 2) // 1 é o buffer

	go func() {
		canal <- 42 // uma gorountine querendo entregar 
		canal <- 43
	}()

	fmt.Println(<-canal) // uma gorountine querendo tirar
	fmt.Println(<-canal)

}