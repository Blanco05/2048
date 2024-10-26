package main

import "fmt"

func main (){
	var orario1, orario2 int
	fmt.Scan(&orario1, &orario2)
	if (orario1 >= 5 && orario2 >= 30) && (orario1 <= 12 && orario2 <= 30){
		fmt.Print("mattino")
	}
}