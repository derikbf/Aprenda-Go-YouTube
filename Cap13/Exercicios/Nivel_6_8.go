package main

import "fmt"

func main() {

	x := retornaumafuncao()
	x()
	
}

func retornaumafuncao() func() {
	return func() {
		fmt.Println("olha eu aqui")
	}
}

// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.