package main

import "fmt"

func main() {
	/*
	 * definire quattro costanti: rettangolo, triangolo, perimetro e area, con queste
	 * gestire la stampa della formula dell'area e del perimetro di triangolo e rettangolo
	*/
	const (
		RETTANGOLO = 1
		TRIANGOLO  = 2
		PERIMETRO  = 1
		AREA       = 2
	)
	var figura, operazione int

	fmt.Println("di che figura si tratta?")
	fmt.Print("rettangolo (1) o triangolo (2)? ")
	fmt.Scan(&figura)
	fmt.Println("cosa vuoi calcolare?")
	fmt.Print("perimetro (1) o area (2)? ")
	fmt.Scan(&operazione)

	if figura == RETTANGOLO {
		if operazione == PERIMETRO {
			fmt.Println("formula = (base + altezza) * 2")
		} else { //operazione == AREA
			fmt.Println("formula = base * altezza")
		}
	} else { //figura == TRIANGOLO
		if operazione == PERIMETRO {
			fmt.Println("formula = lato1 + lato2 + lato3 ")
		} else { //operazione == AREA
			fmt.Println("formula = (base * altezza)/2")
		}
	}
}
/* DOMANDE */
/*
- D1. Quanti diversi messaggi da stampare ci sono?
- sono 4 che si vedono a schermo, una volta fatto eseguire il codice, anche nel caso in cui si sbagli qualcosa
in generale nel codice sono presenti 8 messaggi di stampa

- D2. Sarebbe possibile ridurre il numero di if/else? Perché?
- no non è possibile ridurlo
*/
