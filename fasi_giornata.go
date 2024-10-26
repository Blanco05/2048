package main

import "fmt"

func main(){
	var orario int
	fmt.Println("immetti orario della giornata: ")
	fmt.Scan(&orario)
	if orario >= 0 && orario <= 6{
		fmt.Println("notte")
	} else if orario >= 7 && orario <= 13{
		fmt.Println("mattino")
	} else if orario >= 14 && orario <= 18{
		fmt.Println("pomeriggio")
	} else if orario >= 19 && orario <= 23{
		fmt.Println("sera")
	} else {
		fmt.Println("orario non valido")
	}
}