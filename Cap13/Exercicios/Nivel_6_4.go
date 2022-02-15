package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (odermogorgondatv pessoa) demonstrar() {
		fmt.Println("O nome completo dessa pessoa é:", odermogorgondatv.nome, 
		odermogorgondatv.sobrenome, "\nEssa pessoa tem", odermogorgondatv.idade,
	)
}

func main() {

	aguriazinhadocabeloraspadodatv := pessoa {
		nome: "Doze",
		sobrenome: "Esquisita",
		idade: 12,
	}
	
	aguriazinhadocabeloraspadodatv.demonstrar()	

}
