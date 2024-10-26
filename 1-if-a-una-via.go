//1-if-a-una-via.go

package main

import "fmt"

func main() {
	/*
	 * Scrivere un programma che richieda un voto dell'università e che verifichi
	 * se esso e corretto (compreso tra 0 e 30) in caso contrario stampare 
	 * voto non valido 
	 */

	var voto int

	fmt.Print("voto: ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 30 {
		fmt.Println("voto non valido")
	}
}
