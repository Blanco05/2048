//lezione spiegazione, possibile non usare le flag in questo modo

package main

import(
	"fmt"
)

func pari(x int) bool{
	return  x%2 == 0
}

func main(){
	var x int
	fmt.Println("immetti un numero")
	fmt.Scan(&x)
	if pari(x){
		fmt.Println("pari")
	} else {
		fmt.Println("dispari")
	}
}

//break che interrompe i for, se però ho dei for annidati e il break
//che sta nel for più annidato allora il break mi farà uscire solo dal più interno
//continue forza la prossima esecuzione del ciclo