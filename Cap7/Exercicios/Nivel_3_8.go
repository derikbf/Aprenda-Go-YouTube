package main

import "fmt"

func main() {

	mediaNota := 6

	switch {
		case mediaNota >= 6:
			fmt.Println("Aprovado")
		case mediaNota < 6:
			fmt.Println("Reprovado")
		case mediaNota > 9:
			fmt.Println("Parabéns")
	}
}
