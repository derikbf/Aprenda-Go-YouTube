package main

import "fmt"

func main() {

	qualquercoisa := map[int]string{
		123: "muito legal",
		98 :"menos legal um pouquinho",
		18 :"essa é massa",
		21 :"idade pra ir pra cadeia",
	}

	fmt.Println(qualquercoisa, "\n")

	total := 0

	for key, _ := range qualquercoisa {
		total += key
	}
	fmt.Println(total)

	delete(qualquercoisa, 123)

	fmt.Println(qualquercoisa)
}
