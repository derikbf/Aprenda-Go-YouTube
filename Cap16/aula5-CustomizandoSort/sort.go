package main

import (
	"fmt"
	"sort"
)

func main() {

	ss := []string{"Atirei", "o", "pau", "no", "gato"}
	sort.Strings(ss)
	fmt.Println(ss)
}

