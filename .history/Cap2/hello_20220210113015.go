package main

import "fmt"

func main() {

	x :=  16
	y := "strings"
	z := true

	_, erros := fmt.Println("Hello World", "oi", x)
	fmt.Println(erros)
}
