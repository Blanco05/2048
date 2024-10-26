package main 

import "fmt"

func main(){
	var giorno1, orainizio1, orafine1 int
	var giorno2, orainizio2, orafine2 int

	fmt.Println("Immetti giorno 1")
	fmt.Scan(&giorno1)
	fmt.Println("Immetti ora di inizio 1")
	fmt.Scan(&orainizio1)
	fmt.Println("Immetti ora di fine 1")
	fmt.Scan(&orafine1)
	fmt.Println("Immetti giorno 2")
	fmt.Scan(&giorno2)
	fmt.Println("Immetti ora di inizio 2")
	fmt.Scan(&orainizio2)
	fmt.Println("Immetti ora di fine 2")
	fmt.Scan(&orafine2)
	if giorno1 > 31 || giorno1 < 1{
		fmt.Println("Inserimento sbagliato")
	}
	if giorno2 > 31 || giorno2 < 1{
		fmt.Println("Inserimento sbagliato")
	}
	if orainizio1 > 24 || orainizio1 < 0{
		fmt.Println("Inserimento sbagliato")
	}
	if orainizio2 > 24 || orainizio2 < 0{
		fmt.Println("Inserimento sbagliato")
	}
	if orafine1 > 24 || orafine1 < 0{
		fmt.Println("Inserimento sbagliato")
	}
	if orafine2 > 24 || orafine2 < 0{
		fmt.Println("Inserimento sbagliato")
	}

	if giorno1 == giorno2 && (orainizio1 <= orafine2 && orainizio2 <= orafine1){
		fmt.Println("si sovrappongono")
	} else {
		fmt.Println("non si sovrappongono")
	}
}