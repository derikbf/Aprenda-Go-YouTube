package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

type profissional struct {
	pessoa
	titulo string
	salario int
}


func main() {

	pessoa1 := pessoa{
		nome: "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome: "Maricota",
			idade: 31,
		},
		titulo: "Pizzaola",
		salario: 10000,
	}	

	pessoa3 := pessoa{"Mauricio", 40}
	pessoa4 := profissional{pessoa{"vanderlei", 50}, "Político", 100000}

	fmt.Println("idade: ", pessoa1.idade)
	fmt.Println("Pessoa nome: ", pessoa2.pessoa.nome)
	fmt.Println("nome: ", pessoa2.nome)
	fmt.Println("salario: ", pessoa2.salario)
	fmt.Println("pessoa3: ", pessoa3.nome)
	fmt.Println("pessoa4: ", pessoa4.nome)
}