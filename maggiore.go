// Scrivere un programma Go maggiore.go che legga due interi, li salvi in
// due var min e max nell'ordine in cui li legge; se non sono in ordine crescente,
// li sistemi in modo che min contenga il minore e max il maggiore; 
// infine stampi il contenuto della variabile max.


package main

import "fmt"

func main (){
	var min, max int
	fmt.Print("immetti due variabili: ")
	fmt.Scan(&min, &max)
	


	if min == max{
		fmt.Print("il massimo è: ", max, "\nil minimo è: ", min)
	} else if min < max{
		fmt.Print("il massimo è: ", max, "\nil minimo è: ", min)
	} else if min > max {
		min, max = max, min
		fmt.Print("il massimo è: ", max, "\nil minimo è: ", min)
	}
}