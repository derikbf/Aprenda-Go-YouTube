package main

import "fmt"

func main() {
	
	// ss := [][]string{
	// 	[]string{
	// 		"miu",
	// 		"milton",
	// 		"encher o saco",
	// 	},
	// 	[]string{
	// 		"mimi",
	// 		"martha",
	// 		"pedir comida",
	// 	},
	// 	[]string{
	// 		"meus alunos queridos",
	// 		"que estudamn",
	// 		"exercicios ninja",
	// 	},
	// }
	// for _, v := range ss {
	// 		fmt.Println(v[0])
	// 	for _, item := range v {
	// 		fmt.Println("\t", item)
	// 	}
	// }

	fmt.Println("- Praticando -")
	ssdois := [][]string{
		[]string{
			"Primeira Parte",
			"Parte I - parte 1",
			"Parte I - parte 1",
		},
		[]string{
			"Segunda Parte",
			"Parte II - parte 2",
			"Parte II - parte 2",
		},
		[]string{
			"Terceira Parte",
			"Parte III - parte 3",
			"Parte III - parte 3",
		},
	}
	for _, v := range ssdois {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("-", item)
		}
	}
}
