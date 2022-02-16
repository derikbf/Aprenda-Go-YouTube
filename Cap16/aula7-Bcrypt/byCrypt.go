package main

import "fmt"

func main() {
	
	senha := "20julho1980"
	senhaerrada := "20julho1990"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10) 
	if err != nil { fmt.Println(err) }

	fmt.Println("\n-- Hash -- ")
	fmt.Println(string(sb))

	fmt.Println("\n-- NADA - Senha correta -- ")
  fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println("\n-- Error - senha errada -- ")
  fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))

}