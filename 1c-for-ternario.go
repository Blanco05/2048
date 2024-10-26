package main

import "fmt"

func main() {
    /*
     Scrivere un programma che chieda all'utente 5 numeri e li sommi tra di loro e poi stampi il risultato
    */
    const K = 5
	var n int
	somma := 0
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		somma += n  // s = s + n
	}
	fmt.Println(s)
}
