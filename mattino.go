// Scrivere un programma Go mattino.go che chieda e legga un orario 
// (intero) e stampi "mattino" se l'orario è compreso tra le 6 (incluso)
// e le 13 (escluso).

package main

import "fmt"

func main (){
	var orario int
	fmt.Print("Immetti un orario: ")
	fmt.Scan(&orario)
	if orario >= 6 && orario <= 13{
		fmt.Print("mattino")
	}
}