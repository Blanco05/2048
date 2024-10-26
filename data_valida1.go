package main

import "fmt"

func main() {
	var giorno, mese int
	fmt.Print("Inserisci giorno e mese: ")
	fmt.Scan(&giorno, &mese)

	if mese >= 1 && mese <= 12 && giorno >= 1 && giorno <= 31 {
		fmt.Println("data valida")
	} else {
		fmt.Println("data non valida")
	}
}
