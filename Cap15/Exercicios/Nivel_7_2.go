package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func main() {

	zezinho := pessoa{"Zezinho", "Primodoluizinho", 10}

	fmt.Println(zezinho)
	mudeMe(&zezinho)
	fmt.Println(zezinho)
	
}

func mudeMe (p *pessoa) {
	(*p).nome = "Luizinho"
	p.sobrenome = "Primodozezinho"
}


// Crie um struct "pessoa"
// Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. 
// Essa função deve mudar um valor armazenado no endereço *pessoa.
// 	Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
//  Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
//  "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), 
//  → x.f is shorthand for (*x).f." ←
//  Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
// Crie um valor do tipo pessoa;
// Use a função mudeMe e demonstre o resultado.
