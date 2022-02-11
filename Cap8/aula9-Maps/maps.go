package main

import "fmt"

func main() {

	amigos :=map[string]int{
		"alfredo": 12345678,
		"joana": 9996674,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 444444

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"], "\n\n")

	fmt.Println(amigos["romário"])

	sera, ok := amigos["fantasma"]

	fmt.Println(sera, ok)
}
