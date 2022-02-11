package main

import "fmt"

func main() {
	
	mepe := map[string][]string{
		"dasilva_guriaestranha": []string{
			"desaparecer", "ser assustadora", "raspar a cabeça",
		},
		"Senna_ayrton": []string{
			"andar de jetski", "ser milhonario", "falarrr",
		},
		"pike rob": []string{
			"Ternos loucos", "Roupas estranhas", "wtf",
		},
	}

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " ", h)
		}
	}
}
