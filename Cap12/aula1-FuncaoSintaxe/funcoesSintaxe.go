package main

import "fmt"

func main() {
	
	total, quantos, oi := soma(10, 10, 4, 5, 10, 29)
	fmt.Println("total: ", total, "\nQtd valores: ", quantos, "\n",oi)
}

func soma(x ...int) (int, int, string) {
	soma := 0;
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "bom dia"
}