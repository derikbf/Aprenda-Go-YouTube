package main

import (
	"fmt"
)

// - Chans par, ímpar, quit
// - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
// - Func receive é um select entre os três canais, encerra no quit

func main() {
	par := make(chan int)
	ímpar := make(chan int)
	quit := make(chan bool)

	go mandaNúmeros(par, ímpar, quit)

	receive(par, ímpar, quit)
}

func mandaNúmeros(par, ímpar chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i % 2 == 0 {
			par <- i
		} else {
			ímpar <- i
		}
	}
	close(par)
	close(ímpar)
	quit <- true
}

func receive(par, ímpar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é par.")
		case v := <-ímpar:
			fmt.Println("O número", v, "é ímpar.")
		case <-quit:
			return
		}
	}
}

// - Select é como switch, só que pra canais, e não é sequencial.
// - "A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready." — https://tour.golang.org/concurrency/5
// - Na prática:
// - Exemplo 3:
// - Chans par, ímpar, quit
// - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
// - Func receive é um select entre os três canais, encerra no quit
// - Problema!