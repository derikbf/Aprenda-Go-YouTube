package main

import "fmt"

func main() {
	
	x := []int{42, 43, 44, 45, 46}
	
	y := []int{42, 43, 44, 48, 49, 50, 51}
	x = append(x, y...)

	fmt.Println(x)
}
