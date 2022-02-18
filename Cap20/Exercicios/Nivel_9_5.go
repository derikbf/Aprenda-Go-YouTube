package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32
const quantidadeDeGoroutines = 500

func main() {
	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de Goroutines:\t", quantidadeDeGoroutines,
		"\nTotal do Contador:\t", contador)
}

func criarGoroutines (i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			runtime.Gosched()
			atomic.LoadInt32(&contador)
			fmt.Println(contador)
			wg.Done()
		}()
	}
}

// Utilize atomic para consertar a condição de corrida do exercício #3.