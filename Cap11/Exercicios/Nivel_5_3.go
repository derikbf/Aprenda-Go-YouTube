package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	sabores []string
}

func main() {

	meumapa := make(map[string]pessoa)

	meumapa["Pimentão"] = pessoa {
		nome: "Renata",
		sobrenome: "Pimentel",
		sabores: []string{"pistache", "morango", "baunilha"}}

	meumapa["da prucia"] = pessoa {"Frederico", "da prucia", 
		[]string{"sabao em pó", "chocolate", "agua salgada"}}
	
	for _, v := range meumapa {
		fmt.Println(v)
	}
}
