package main

import "fmt"

func main() {

	quemestanoescritoriohoje := "derik"

	switch quemestanoescritoriohoje {
		case "zezinho":
			fmt.Println("hoje quem está no escritorio é o Zezinho")
			fallthrough
		case "marquinhos":
			fmt.Println("sempre que o zezinho vem, o marquinhos vem também")
		case "joana":
			fmt.Println("hoje quem tá no escritorio é a Joana")
		case "Maria":
			fmt.Println("hoje quem tá no escritorio é a Maria")
		default:
			fmt.Println("o escritório tá vazio")
	}

	quemtanafirma := "zezinho"

	switch quemtanafirma {
		case "zezinho", "maria":
			fmt.Println("hoje quem tá na firma é o time 1")
		case "joana", "pedinho":
			fmt.Println("hoje quem tá na firma é o time 2")
	}
}
