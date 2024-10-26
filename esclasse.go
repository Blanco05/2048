package main

import ("fmt")

func main() {
	numprima, denprima := 0,0
	fmt.Println("Immetti due numeri per la prima frazione")
	fmt.Scan(&numprima, &denprima)

	numseconda, denseconda := 0,0
	fmt.Println("Immetti due numeri per la prima frazione")
	fmt.Scan(&numseconda, &denseconda)

	frazione1 := float64(numprima) / float64(denprima)
	frazione2 := float64(numseconda) / float64(denseconda)

	if frazione1 < frazione2{
		fmt.Println("la prima frazione è minore della seconda")
	} else {
		fmt.Println("la prima frazione è maggiore della seconda")
	}

}