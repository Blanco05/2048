package main

import "fmt"

func main() {
	var giorno, mese int
	fmt.Print("Inserisci giorno e mese: ")
	fmt.Scan(&giorno, &mese)

	if (mese == 1 || mese == 3 || mese == 5 || mese == 7 || mese == 8 || mese == 10 || mese == 12) && giorno >= 1 && giorno <= 31 {
		fmt.Println("data valida")
	} else if (mese == 4 || mese == 6 || mese == 9 || mese == 11) && giorno >= 1 && giorno <= 30 {
		fmt.Println("data valida")
	} else if mese == 2 && giorno >= 1 && giorno <= 28 {
		fmt.Println("data valida")
	} else {
		fmt.Println("data non valida")
	}
}
