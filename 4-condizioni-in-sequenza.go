package main

import "fmt"

func main() {
	/* 	 
	 * Scrivere un programma che richieda un intero, voto, e che verifichi
	 * se esso p maggiore di 90 e scrivi A, maggiore di 80 scriva B,
	 * maggiore di 70 scriva C, maggiore di 60 scriva D, sotto 60 scriva
	 * F 
	 */

	var voto int
	fmt.Print("voto? ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 100 {
		fmt.Println("voto non valido")
		return //interrompe l'esecuzione del programma
	}

	if voto >= 90 {
		fmt.Println("A")
	} else if voto >= 80 {
		fmt.Println("B")
	} else if voto >= 70 {
		fmt.Println("C")
	} else if voto >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
/* DOMANDE */
/*
- D1. Per quali valori è vera la condizione del primo if? (quello subito dopo la Scan)
- per i valori sopra o uguali a 90

- D2. Per quali valori viene eseguita la seguente istruzione?
    fmt.Println("B")
- viene stampata solo se il voto è compreso tra 80 incluso e 90 escluso

- D3. Giustificate la vostra risposta alla domanda D2
- perchè entra nel primo if verifica che non è maggiore o uguale a 90
  e quindi entra nel secondo e ci rimane solo se è maggiore o uguale a 80

- D4. Il programma è impostato correttamente o sarebbe stato (più) corretto scrivere invece così?
 } else if voto >= 80 && voto < 90 {
- è impostato correttamente

- D5. Perché?
- perchè entra nel secondo if se e solo se ha verificato che il voto non è maggiore o uguale di 90
*/
