package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("SO:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
}

// Crie um programa que demonstra seu OS e ARCH.
// Rode-o com os seguintes comandos:
//  go run
//  go build
//  go install