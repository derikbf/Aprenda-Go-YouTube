package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Gorountines1:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 10

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)


	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			v := contador
			runtime.Gosched() // parecido com o time.Sleep()
			v++
			contador = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Gorountines2:", runtime.NumGoroutine())
	fmt.Println("valor final:", contador)

}
