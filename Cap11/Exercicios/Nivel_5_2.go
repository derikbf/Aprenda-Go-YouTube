package main

import "fmt"

type veiculo struct {
	portas int
	cor string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatros bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

		carraoDoTio := sedan{veiculo{4, "abobora"}, true}
		fubicaDoVô := caminhonete{
			veiculo: veiculo{
				portas: 8,
				cor: "ferrugagem",
			},
			tracaoNasQuatros: false,
		}

		fmt.Println(carraoDoTio)
		fmt.Println(fubicaDoVô)

		fmt.Println(carraoDoTio.portas)
		fmt.Println(fubicaDoVô.cor)
}
