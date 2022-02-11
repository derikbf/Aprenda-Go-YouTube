package main

import "fmt"


func main() {

	array := [5]int{1, 2, 3, 4, 5}
		fmt.Println("-- array --")	
		fmt.Println(array)
		fmt.Println("\n-- slice --")
		slice := []int{1, 2, 3, 4, 5}
		fmt.Println(slice)
		fmt.Println(slice[3])

		fmt.Println("\n-- slice 2 --")
		slice[3] = 348756
		fmt.Println(slice[3])
}