package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

type dentista struct {
	pessoa 
	dentesarrancados int
	salario float64
}

type arquiteto struct {
	pessoa
	tipodeconstrucao string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é ", x.nome, 
	"e eu já arranquei", x.dentesarrancados, "e ouve só: Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é ", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
		switch g.(type) {
		case dentista:
			fmt.Println("Eu ganho:", g.(dentista).salario)

		case arquiteto:
			fmt.Println("Eu construo:", g.(arquiteto).tipodeconstrucao)
		}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome: "Mister Dente",
			sobrenome: "da Silva",
			idade: 50,
		},
		dentesarrancados: 8973,
		salario: 8950,
	}
	mrpredio := arquiteto{
		pessoa: pessoa{
			nome: "Mister Prédio",
			sobrenome: "Casas Bahia",
			idade: 55,
		},
		tipodeconstrucao: "hospicios",
		tamanhodaloucura: "demais",
	}

	serhumano(mrdente)
	fmt.Println("----")
	serhumano(mrpredio)
	// mrdente.oibomdia()
	// mrpredio.oibomdia()

}