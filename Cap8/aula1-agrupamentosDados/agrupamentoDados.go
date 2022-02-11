package main

import "fmt"

var x [5]int

func main() {

	x[0] = 1
	x[1] = 10
	x[2] = 100
	x[3] = 1000
	x[4] = 10000

	fmt.Println(x[0], x[1], x[2])
	fmt.Println("----")
	fmt.Println(x)
	fmt.Println("----")
	fmt.Println("%T", x)
	fmt.Println("----")
	fmt.Println(len(x))
}