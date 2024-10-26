package main

import(
	"fmt"
)

func lettura(str string) int{
	conta := 0
	for _, c:=range str{
		if c == 'A' || c == 'a'{
			conta++;
		} 
	}
	return conta
}

func main(){
	var str string
	str = "ciààao"
	numero := lettura(str)
	fmt.Println(numero)
}

// un programma che data una stringa e un intero lui verifichi quanti caratteri ci sono tra la barra iesima e la barra i + 1 esima
// ciao/come/stai i = 1  i = 2 -> stampa 4


