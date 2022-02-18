package main

import "fmt"

func main() {

	canal := make(chan int)

	go func() {
		canal <- 42 // uma gorountine querendo entregar 
	}()

	fmt.Println(<-canal) // uma gorountine querendo tirar

}