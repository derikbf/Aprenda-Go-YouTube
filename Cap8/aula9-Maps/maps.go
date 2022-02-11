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

	if sera, ok := amigos["fantasma"]; !ok {
		fmt.Println("Não tem!")
	} else {
		fmt.Println(sera, ok)
	}
}
