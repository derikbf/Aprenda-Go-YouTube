package main

import "fmt"

func main() {
	
	slice := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 10}

	fmt.Printf("%T\n", slice)

	fmt.Println(slice[0:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	

}