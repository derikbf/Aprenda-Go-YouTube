package main

import "fmt"

func main() {
	s := "ascii, é0â"

	for _, v := range s {
		fmt.Printf("%v- %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v- %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}
