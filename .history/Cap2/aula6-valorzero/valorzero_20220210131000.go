package main

import "fmt"

// var x int

func main() {

	x = 10
	fmt.Printf("%v, %T\n", x, x)

	x = 20
	fmt.Printf("%v, %T", x, x)
}
