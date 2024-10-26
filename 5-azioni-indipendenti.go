package main

import "fmt"

func main() {
	/* 
	 * Scrivere un programma che richieda un numero intero all'utente e che
	 * stampi Fizz se esso è divisibile per 3, Buzz se esso è divisibile per
	 * 5
	 */

	var num int
	fmt.Print("numero? ")
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("Fizz")
	}
	if num%5 == 0 {
		fmt.Println("Buzz")
	}
}
/* DOMANDE */
/*
- D1. Se al posto del secondo if ci fosse un else if (legato al primo if), il programma si comporterebbe nello stesso modo? Dareste le stesse specifiche?
- il programma si comporterebbe allo stesso modo, solo nel caso in cui il numero non fosse divisibile per 3
  in caso contrario non entrerebbe nel secondo if, quindi cambierei le specifiche scrivendo che se il numero non è divisibile per tre allora deve uscire dagli if
  se in caso contrario deve verificare se è anche divisibile per 5

- D2. Perché? Spiegate la vostra risposta alla domanda D1
- se io creo un programma che entra in un primo if e quell'if è verificato non erntrerà nel secondo if

*/
