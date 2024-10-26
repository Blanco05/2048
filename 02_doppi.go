package main

import(
	"fmt"
)

func main(){
	const K = 5
	var n int
	for i:=0; i<K; i++{
		fmt.Println("Immetti un", i, "numero")
		fmt.Scan(&n)
		fmt.Print(n*n, "\n")
	}
}