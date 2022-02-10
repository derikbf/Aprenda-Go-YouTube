package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := "e"
	b := "é"
	
	fmt.Println("---------")
	fmt.Println("%v, %v, %v", a, b)

	d := []byte(a)
	e := []byte(b)
	
	fmt.Println("---------")
	fmt.Printf("%v, %v, %v", d, e)

	x := 10
	xf := 10.0
	
	fmt.Println("---------")
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", xf, xf)

	fmt.Println(" -- SO --")
	fmt.Println(runtime.GOOS, "\n")
	fmt.Println("-- PROCESSADOR --")
	fmt.Println(runtime.GOARCH)
}
