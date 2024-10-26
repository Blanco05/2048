package main 

import "fmt"

func main (){
	var mionumero = 7
	var indovinare int
	fmt.Println("immetti un numero tra 1-10: ")
	fmt.Scan(&indovinare)
	if indovinare > 10 || indovinare < 0{
		fmt.Println("Valore non valido")
	} else {
		if indovinare == mionumero {
			fmt.Println("Hai indovinato!")
		} else {
			fmt.Println("Non hai indovinato")
		}
	}
}