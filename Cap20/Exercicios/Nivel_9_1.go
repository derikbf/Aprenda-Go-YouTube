package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	novagorountines(5)
	wg.Wait()
}

func novagorountines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Eu sou a goroutine número", i+1)
			wg.Done()
		}(x)
	}
}

// - Alem da goroutine principal, crie duas outras goroutines.
// - Cada goroutine adicional devem fazer um print separado.
// - Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.