package main

import "fmt"

func main() {

	fmt.Println("Esse é o nosso programa do ex de compilação
		cruzada. foi compilado num linux/amd64, e agora está
		rodando num sistema:", runtme.GOARCH, runtime.GOOS)
	
}