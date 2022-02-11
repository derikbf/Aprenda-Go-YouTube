package main

import "fmt"


func main() {

	slice := []string{"banana", "maça", "jaca", "pêssego"}
	for i, v := range slice {
		fmt.Println("No íncide", i, "temos o valor", v)
	}
	
	//slice[4] = "melancia"
	slice = append(slice, "melancia")

	for _, v := range slice {
		fmt.Printf("Um dos valores desse slice é %v.\n", v)
	}

}