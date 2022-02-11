package main

import "fmt"

func main() {
	
	anoemqueeunasci := 2020
	anoateoqualeuquerocontar := 2025

	for {
		if (anoemqueeunasci > anoateoqualeuquerocontar) {
			break
		}
		fmt.Println(anoemqueeunasci)
		anoemqueeunasci++
	}
}
