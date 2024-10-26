package main

import(
	"fmt"
)

/*
 * la funzione calcola il coefficiente angolare
 */

func calcola(x1 float64, y1 float64, x2 float64, y2 float64) float64{
	return (y2-y1)/(x2-x1)
}

func main(){
	var x1, y1, y2, x2 float64
	fmt.Println("immetti x e y del primo punto: ")
	fmt.Scan(&x1, &y1)
	fmt.Println("immetti x e y del secondo punto: ")
	fmt.Scan(&x2, &y2)
	fmt.Println("il coeff. angolare è: ", calcola(x1, y1, y2, x2))
}