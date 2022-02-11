package main

import "fmt"

var x interface{}

func main() {

	x = 10
	
	switch x.(type) {
		case int:
			fmt.Println("é um int")
		case bool:
			fmt.Println("é um bool")
		case string:
			fmt.Println("é um string")
		case float64:
			fmt.Println("é um float")
		default:
			fmt.Println("nada")
	}
}
