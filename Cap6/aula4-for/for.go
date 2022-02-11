package main

import "fmt"

func main() {

	x := 0
	
	for {
		if x < 10 { x ++ 
			fmt.Println(x, "x é menor que 10")	
		}	else { 
			fmt.Println("X é maior que 10")
			break
		}
	}
}