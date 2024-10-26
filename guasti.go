package main

import "fmt"

func main(){
	var guasto1, guasto2, guasto3 int
	fmt.Println("componente 1, guasto rilevato? (0 per ok)")
	fmt.Scan(&guasto1)
	fmt.Println("componente 2, guasto rilevato? (0 per ok)")
	fmt.Scan(&guasto2)
	fmt.Println("componente 3, guasto rilevato? (0 per ok)")
	fmt.Scan(&guasto3)
	fmt.Println("guasti rilevati a:")
	if guasto1 == 0 {
		fmt.Println("caricamento acqua")
	}
	if guasto2 == 0{
		fmt.Println("scarico acqua")
	}
	if guasto3 == 0{
		fmt.Println("sistema di riscaldamento")
	}
}