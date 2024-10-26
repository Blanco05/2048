package main

import "fmt"

func main(){
	var x int
	fmt.Println("immetti un numero: ")
	fmt.Scan(&x)
	if x%10 == 0{
		fmt.Println("divisibile per 10")
	} 
	if x > 0{
		fmt.Println("il numero è positivo")
	}
	if x > 0 && x%10 == 0{
		fmt.Println("positivo divisibile per 10")
	}
}

//i diversi if che mi servono sono in in colonna, if a una via, tre del primo tipo