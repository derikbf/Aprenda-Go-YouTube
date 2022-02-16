package main

import (
	"fmt"
	"sort"
)

func main() {

	ss := []string{"A", "C", "B", "F", "D"}
	sort.Strings(ss)
	fmt.Println(ss)

	ssi := []int{1, 4, 6, 8, 9}
	sort.Ints(ssi)
	fmt.Println(ssi)
}

