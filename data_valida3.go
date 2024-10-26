package main

import "fmt"

func main() {
	var giorno, mese, anno int
	fmt.Print("Inserisci giorno, mese e anno: ")
	fmt.Scan(&giorno, &mese, &anno)

	isBisestile := func(anno int) bool {
		return (anno%4 == 0 && anno%100 != 0) || (anno%400 == 0)
	}

	dataValida := false
	if (mese == 1 || mese == 3 || mese == 5 || mese == 7 || mese == 8 || mese == 10 || mese == 12) && giorno >= 1 && giorno <= 31 {
		dataValida = true
	} else if (mese == 4 || mese == 6 || mese == 9 || mese == 11) && giorno >= 1 && giorno <= 30 {
		dataValida = true
	} else if mese == 2 {
		if isBisestile(anno) && giorno >= 1 && giorno <= 29 {
			dataValida = true
		} else if giorno >= 1 && giorno <= 28 {
			dataValida = true
		}
	}

	if dataValida {
		fmt.Println("data valida")
	} else {
		fmt.Println("data non valida")
	}
}
