package main

import "fmt"

func main() {
	fmt.Println(retornaUmInt())
	fmt.Println(retornaUmInteUmaString())
}

func retornaUmInt() int {
	return 10
}

func retornaUmInteUmaString() (int, string) {
	return 20, "vinte"
}