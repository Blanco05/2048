package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Print("immetti quanti quadrati vuoi?")
	fmt.Scan(&n)
	for i:=0; i<n; i++{
		fmt.Println(i*i)
	}
}