package main

import "fmt"

func main() {
	x :=  10
	a :=  10 < 20
	// y := "olá bom dia"
	// x := true

	fmt.Println(a, "\n")
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	
	x, z := 20, 30
	fmt.Printf("y: %v, %T\n\n", x, x)
	fmt.Printf("y: %v, %T\n\n", z, z)

}
