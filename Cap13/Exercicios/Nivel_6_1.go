package main

import "fmt"

func main() {
	fmt.Println(retornaUmInt())
	fmt.Println(retornaUmInteUmaString())
}

func retornaUmInt() int {
	return 10
}

func retornaUmInteUmaString() (int, string) {
	return 20, "vinte"
}

// - Exercício:
//     - Crie uma função que retorne um int
//     - Crie outra função que retorne um int e uma string
//     - Chame as duas funções
//     - Demonstre seus resultados