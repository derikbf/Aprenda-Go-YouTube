package main

import "fmt"

const (
	a = iota
	_ = iota
	c = iota
	x = iota
	y = iota
	z = iota
)

func main() {

	fmt.Println(a, c, x, y, z)

}
