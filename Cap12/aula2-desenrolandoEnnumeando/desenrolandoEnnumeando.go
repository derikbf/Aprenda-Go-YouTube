package main

import "fmt"

func main() {
	
	si := []int{10, 1,5,8,9}
	
	total := soma(si...)
	fmt.Println("total: ", total)
}

func soma(x ...int) int {
	soma := 0;
	for _, v := range x {
		soma += v
	}
	return soma
}