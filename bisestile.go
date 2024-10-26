package main

import "fmt"

func main() {
	var anno int
	fmt.Print("Inserisci l'anno: ")
	fmt.Scan(&anno)

	if (anno%4 == 0 && anno%100 != 0) || (anno%400 == 0) {
		fmt.Println("bisestile")
	} else {
		fmt.Println("non bisestile")
	}
}
