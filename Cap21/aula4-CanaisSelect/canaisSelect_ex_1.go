package main

import (
	"fmt"
)

func main() {

	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)
	
	for i := 0; i < x; i++ {
		select {
			case v := <-a:
				fmt.Println("Canal A:", v)			
			case v := <-b:
				fmt.Println("Canal B:", v)
		}
	}
}

// - Select é como switch, só que pra canais, e não é sequencial.
// - "A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready." — https://tour.golang.org/concurrency/5
// - Na prática:
//     - Exemplo 1:
//         - Duas go funcs enviando X/2 numeros cada uma pra um canal
//         - For loop X valores, select case ←x
