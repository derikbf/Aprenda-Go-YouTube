package main

import (
	"encoding/json"
	"os"
)

type pessoa struct {
	Nome string
	Sobrenome string
	Idade int
	Profissao string
	Contabancaria float64
}

// type informacoes struct {
// 	Nome string `json:"Nome"`
// 	Sobrenome string `json:"Sobrenome"`
// 	Idade int `json:"Idade"`
// 	Profissao string `json:"Profissao"`
// 	Contabancaria float64 `json:"Contabancaria"`
// }

func main() {
	jamesbond := pessoa{
		Nome: "James",
		Sobrenome: "Bond",
		Idade: 40,
		Profissao: "Agente Secreto",
		Contabancaria: 100.5,
	}	

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesbond)

	// sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":100.5}`)
	
	// var jamesbond informacoes 
	// err := json.Unmarshal(sb, &jamesbond)	
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// fmt.Println(jamesbond)
	// fmt.Println(jamesbond.Profissao)
}

// {"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":100.5}
