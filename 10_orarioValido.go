package main

import(
	"fmt"
)

func main(){
	var h, min int
	for{
		fmt.Print("Ore e Minuti: ")
		fmt.Scan(&h, &min)
		if (h < 23 && h > 0) && (min > 0 && min < 59){
			fmt.Println("valori validi")
			break
		}
 	}
}