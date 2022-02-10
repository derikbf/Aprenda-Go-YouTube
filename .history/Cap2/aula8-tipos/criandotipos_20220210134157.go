package main

import "fmt"

type hotdog int 

var b hotdog = 10

func main() {

	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", b)
}
