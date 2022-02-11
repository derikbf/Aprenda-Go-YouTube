package main

import "fmt"

func main() {

	sabores := []string{"peperoni", "mozzarela", "quatroQ", "abacaxi", "marg"}
	
	sabores = append(sabores[:2], sabores[4:]...)

	fmt.Println(sabores[0], sabores[1], sabores[2])

	fmt.Println(sabores)

}
