package main

import "fmt"

func main() {

	for i := 0; i < 20; i++ {
		if i % 2 != 0 {
			fmt.Println(i, "é ímpar")
		} else {
			fmt.Println(i, "é par")
		}
		// fmt.Println("Hello")
	}
}
