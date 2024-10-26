package main

import "fmt"

func main() {
	var i int
	var prezzo, litri float64
	benzina := 1.78
	diesel := 1.98
	alcol := 1.2
	cherosene := 1.1

	fmt.Println("Immetti quanti litri: ")
	fmt.Scan(&litri)
	fmt.Println("Scegli il tipo di carburante:\n0 - benzina (1.78)\n1 - diesel (1.98)\n2 - alcol (1.2)\n3 - cherosene (1.1)")
	fmt.Scan(&i)

	if i == 0 {
		prezzo = litri * benzina
	} else if i == 1 {
		prezzo = litri * diesel
	} else if i == 2 {
		prezzo = litri * alcol
	} else if i == 3 {
		prezzo = litri * cherosene
	} else {
		fmt.Println("Opzione non valida")
		return
	}

	fmt.Printf("Il prezzo è: %.2f\n", prezzo)
}
