package main

import (
	"fmt"
	"math"
)

// Aqui está uma interface básica para formas geométricas.
type geometry interface {
	area() float64
	perim() float64
}

// Para nosso exemplo, implementaremos essa interface em recte circletipos.
type rect struct {
	width, height float64
}
type circle struct { 
	radius float64
}

// Para implementar uma interface em Go, basta implementar todos os métodos
// da interface. Aqui nós implementamos geometry em rects.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// A implementação para circles.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Se uma variável tiver um tipo de interface, podemos chamar métodos que 
// estão na interface nomeada. Aqui está uma measurefunção genérica 
// aproveitando isso para trabalhar em qualquer geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}


// Os tipos struct circlee struct implementam a interface para que possamos
// usar instâncias dessas structs como argumentos para .rectgeometrymeasure