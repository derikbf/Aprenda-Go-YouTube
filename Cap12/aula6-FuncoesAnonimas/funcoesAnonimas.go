package main

import "fmt"


func main() {

	x := 25

	func(x int) {
		fmt.Println(x, "vezes 10 é:")
		fmt.Println("total:", x * 10)
	}(x)
	
	// fmt.Println("Hello")
}