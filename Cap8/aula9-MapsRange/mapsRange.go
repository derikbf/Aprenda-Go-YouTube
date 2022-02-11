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

	for key, value := range qualquercoisa {
		fmt.Println(key, value)
	}
}
