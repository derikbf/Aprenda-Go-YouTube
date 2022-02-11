package main

import "fmt"


func main() {

	x := 9

	if (x % 2 == 0 && x % 3 == 0) {
		fmt.Println("é multiplo de dois e treis")
	} 
	
	if x%2 == 0 || x%3 == 0 {
		fmt.Println("é multiplo de dois ou de treis")
	}
}
