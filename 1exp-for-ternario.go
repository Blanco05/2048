package main

import "fmt"

func main() {
    /*
     definita una costante BASE = 8, chiedere un int, se questo int diviso la base è minore della lunghezza del numero in float allora stampare una barra, separata da dei \t. 
     un secondo for che ha lastessa condizione del primo, ma stampa il risultato tra la divisione di i e BASE e il resto tra i e BASE
    */

	const BASE = 8
	var n int

	fmt.Print("un int: ")
	fmt.Scan(&n)

	length := float64(n)
	for i := 0; float64(i)/BASE < length; i++ { // i < n*BASE
		fmt.Print("|\t")
	}
	fmt.Println()

	for i := 0; float64(i)/BASE < length; i++ {
		fmt.Print(i/BASE, i%BASE, "\t")
	}
	fmt.Println()
}
