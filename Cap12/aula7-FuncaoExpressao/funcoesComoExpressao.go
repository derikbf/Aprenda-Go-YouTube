package main

import "fmt"

func main(){

	x := 12

	y := func(x int) int {
		return x * 20
	}

	fmt.Println(x, "vezes 20 é:", y(x))
}