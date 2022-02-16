package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome string
	potencia int
	consumo int
}
// -------- Potencia 
type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int {
	return len(x)
}

func (x ordenarPorPotencia) Less(i, j int) bool {
	return x[i].potencia < x[j].potencia
}

func (a ordenarPorPotencia) Swap(i, j int) {
	a[i], a[j] = a[j], a[i] 
}

// -------- Consumo
type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int {
	return len(x)
}

func (x ordenarPorConsumo) Less(i, j int) bool {
	return x[i].consumo > x[j].consumo
}

func (a ordenarPorConsumo) Swap(i, j int) {
	a[i], a[j] = a[j], a[i] 
}

// Por Dono do Posto 
type ordenarPorLucroProDonoDoPosto []carro

func (x ordenarPorLucroProDonoDoPosto) Len() int {
	return len(x)
}

func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool {
	return x[i].consumo < x[j].consumo
}

func (a ordenarPorLucroProDonoDoPosto) Swap(i, j int) {
	a[i], a[j] = a[j], a[i] 
}

func main() {

	carros := []carro{{"Chevete", 50, 8},
										{"Porsche", 300, 5},
										{"Fusca", 20, 30},
										}

	fmt.Println(carros)
	fmt.Println("\n-- Ordenar Por Potencia --")
	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println(carros)
	
	fmt.Println("\n-- Ordenar Por Consumo --")
	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println(carros)

	fmt.Println("\n-- Ordenar Por Lucro Pro Dono do Posto --")
	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))
	fmt.Println(carros)
}
