package main

import "fmt"

func main() {

	for horas := 0; horas <= 24; horas++ {
		fmt.Println("Hora: ", horas)
	
		for x := 0; x < 60; x++ {
			fmt.Print(" ", x)
		}
		fmt.Println("\n")
	}
}