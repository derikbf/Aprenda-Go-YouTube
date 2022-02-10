package main

import "fmt"

var y = 10

func main() {
	
	z := 20
	qualquercoisa(y)
}

func qualquercoisa (x int) {
	fmt.Println(y)
	fmt.Println(x)
}
