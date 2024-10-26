package main

import(
	"fmt"
)

func main(){
	var n, cifra int
	fmt.Print("immeti numero: ")
	fmt.Scan(&n)
	posizione := 1
	for n>0{
		cifra = n%10
		if cifra%2 == 0{
			fmt.Println(posizione)
           	break
	       }
	       n = n / 10
	       posizione++
	       if n == 1{
	       		fmt.Println(-1)
	       }
   }
   
}