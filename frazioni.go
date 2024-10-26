package main

import "fmt"

func main (){
	var x1, x2, y1, y2 int
	fmt.Print("immetti la prima frazione: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("immetti la seconda frazione: ")
	fmt.Scan(&x2, &y2)
	// x1/y1 = x2/y2
	// x1*y2 = x2*y1

	if x1*y2 == x2*y1{
		fmt.Print("equilaventi")
	} else {
		fmt.Print("non equilaventi")
	}
}