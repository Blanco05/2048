package main

import "fmt"

func main() {

/* Qui sotto trovate quattro commenti (a., b., c., d.) e poi quattro cicli for.
Per ciascun ciclo for scegliete il commento corrispondente a quello che fa il ciclo for e spostatelo (taglia e incolla) al posto della riga "//...". */

// a. stampa i numeri tra 0 e n (escluso) espressi in base 3
// b. stampa le potenze di 2 tra 0 e n (escluso)
// c. stampa gli interi quadrati perfetti tra 0 e n (escluso)
// d. calcola la somma dei primi n valori non negativi letti

	// ciclo 1
	//d. calcola la somma dei primi n valori non negativi letti
	var value, sum float64
	for i := 0; i < n; i++ {
    	fmt.scan(&value)
    	if value < 0 { 
			i--
   		} else {
        	sum += value
   		}
	}

	// ciclo 2
	//b. stampa le potenze di 2 tra 0 e n (escluso)
	for i := 0; i < n; i *= 2 {
		fmt.Print(i, "\t")
	}

	// ciclo 3
	//a. stampa i numeri tra 0 e n (escluso) espressi in base 3
	const BASE = 3
	n := BASE*BASE
	for i := 0; i < n; i++ {
		fmt.Print(i, "\t", i/BASE, i%BASE, "\n")
	}

	// ciclo 4
	//c. stampa gli interi quadrati perfetti tra 0 e n (escluso)
	for i := 0; i*i < n; i++ {
		fmt.Print(i*i, "\t")
	}
}
