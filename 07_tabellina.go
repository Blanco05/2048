package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Println("immetti un numero di cui vuoi la tabellina: ")
	fmt.Scan(&n)
	for i:=1; i<11; i++{
		fmt.Print(i*n, "\t")
	}
}