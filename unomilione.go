//scrivere una funzione che calcola quanti zeri compaiono nella scrittura 
//di un numero e poi usatela per calcolare quanti zeri dovete scrivere se 
//scrivete tutti i numeri da 1 a 1milione
package main

import(
	"fmt"
)

func divisore(x int, y int) bool{
	return x%y == 0
}

func main(){
	var x,y int
	fmt.Println("Immetti due valori: ")
	fmt.Scan(&x,&y)
	if divisore(x,y) {
		fmt.Println("Sono divisibili")
	} else {
		fmt.Println("Non sono divisibili")
	}
}