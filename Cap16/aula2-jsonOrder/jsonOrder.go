package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome string
	Sobrenome string
	Idade int
	Profissao string
	Contabancaria float64
}

func main() {
	jamesbond := pessoa{
		Nome: "James",
		Sobrenome: "Bond",
		Idade: 40,
		Profissao: "Agente Secreto",
		Contabancaria: 100.50,
	}

	darthvader := pessoa{"Anakin", "Skywalker", 50, "Vilao Intergaláctico", 50.98}

	j, err := json.Marshal(jamesbond)
	if err != nil {
		fmt.Println(err)
	}
	
	d, err := json.Marshal(darthvader)
	if err != nil {
		fmt.Println(err)
	}	

	fmt.Println(string(j))
	fmt.Println(string(d))
}
