package main

import(
	"fmt"
)

func main(){
	for i:=0; i<10; i++{
		for j:=0; j<i; j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}