package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	quit := make(chan int)
	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
}

func recebeQuit (canal chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido:", <-canal)
	}
	quit <- 0
}

func enviaPraCanal (canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
			case canal <- qualquercoisa:
				qualquercoisa++
			case <-quit:
				return
		}
	}
}

// - Select é como switch, só que pra canais, e não é sequencial.
// - "A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready." — https://tour.golang.org/concurrency/5
// - Na prática:
// - Exemplo 2:
// - Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
// - Func 2 for infinito, select: case envia pra canal, case recebe de quit
