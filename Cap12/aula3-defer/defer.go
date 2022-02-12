package main

import "fmt"

func main() {
	
	defer fmt.Println("com defer, veio primeiro - sairá por último")
	defer fmt.Println("1 - sairá penultimo")
	defer fmt.Println("2 - sairá antepenultimo")
	fmt.Println("sem defer, primeiro")
}
