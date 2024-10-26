package main

import "fmt"

func main() {
	/*
	 * Scrivere un programma che prenda un intero, controlli se è maggiore
	 * di zero, uguale a zero o minore di zero e stampi, in ordine: positivo
	 * nullo o negativo
	 */

	var num int

	fmt.Print("un int: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("positivo")
	} else if num == 0 {
		fmt.Println("nullo")
	} else { // < 0
		fmt.Println("negativo")
	}
}
