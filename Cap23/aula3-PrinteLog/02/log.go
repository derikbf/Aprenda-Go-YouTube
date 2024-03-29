package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}

/*
Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
*/

// log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.