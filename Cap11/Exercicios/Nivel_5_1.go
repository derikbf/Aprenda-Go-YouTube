package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	sabores []string
}

func main() {
	pessoa1 := pessoa {
		nome: "Renata",
		sobrenome: "Pimentel",
		sabores: []string{"pistache", "morango", "baunilha"}}

	pessoa2 := pessoa {"Frederico", "da prucia", 
		[]string{"sabao em pó", "chocolate", "agua salgada"}}
	
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}
