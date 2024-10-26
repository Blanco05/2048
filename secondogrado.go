//dati A,B,C decidere se l'equazione ax^2 + bx + c = 0

package main

import ("fmt")

func main (){
	a,b,c := 0,0,0
	fmt.Println("Immetti i tre indici della tua equazione: ")
	fmt.Scan(&a, &b, &c)

	delta := b*b - 4*(a)*(c)
	fmt.Println(delta)
	if delta > 0{
		fmt.Println("la funzione ha due soluzioni distinte")
	} else {
		fmt.Println("la funzione non ha due soluzioni distinte")
	}

}
