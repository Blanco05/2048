//determina se uno è un divisore di un altro con variabili booleane

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