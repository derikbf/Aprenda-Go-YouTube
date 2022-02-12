package main

import "fmt"

type carro struct {
	nome string
	marca string
	cor string
	ano int
	qtdPortas int
}

func (c carro) carroanda() {
	fmt.Println("Tentando ligar o carro: do ano", c.ano)
	fmt.Println("Nome:", c.nome, "\nMarca:", c.marca, 
	"\nCor:", c.cor, "\nQtd Portas:", c.qtdPortas, "\nbrunm..brunm..LIGOU....RUN")
}

func main () {
	fusca := carro{"Fusca", "Fusquinha", "Amarelo", 1985, 2}
	fusca.carroanda()
}

// resolução da professora 
// package main

// import "fmt"

//  type pessoa struct {
// 	 nome string
// 	 idade int
//  }

//  func (p pessoa) oibomdia() { 
// 	 // oibomdia é uma função especifica para valores daquele tipo
// 	 fmt.Println("Nome: ", p.nome, "de", p.idade, "anos", "diz bom dia!")
//  }

// func main() {
// 	mauricio := pessoa{"Mauricio", 30}

// 	mauricio.oibomdia()
// }
