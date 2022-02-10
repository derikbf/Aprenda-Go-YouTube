package main

import "fmt"

type hotdog int
var x hotdog

func main() {
	fmt.Println(" Exercicio nível #1 - 4")
	
	fmt.Println("valor inicial de x")
	fmt.Println(x)
	fmt.Printf("tipo de x é: %T\n", x)
	x = 42 
	fmt.Println("novo valor de x")
	fmt.Println(x)
	
}
