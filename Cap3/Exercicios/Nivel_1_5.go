package main

import "fmt"

type hotdog int
var x hotdog
var y int

func main() {
	fmt.Println(" Exercicio nível #1 - 5")
		
	x = 42
	y = int(x)
	fmt.Printf("o tipo de x é: %T\n", x)
	fmt.Printf("o tipo de y é: %T\n", y)
	
}
