package main

import "fmt"

//Scrivi un programma che verifichi se un numero 
//inserito dall'utente è pari o dispari.

func main(){
	var x int
	fmt.Println("Inserisci un numero: ")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Println("Il numero e' pari")
	} else {
		fmt.Println("Il numero e' dispari")
	}


}



