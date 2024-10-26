package main

import "fmt"

func main() {
	const aliquotaBassa = 0.10
	const aliquotaAlta = 0.24
	const sogliaNonConiugato = 32000
	const sogliaConiugato = 64000

	var reddito int
	var coniugato bool
	var tasse float64

	fmt.Print("reddito? ")
	fmt.Scan(&reddito)
	fmt.Print("coniugato? (true/false) ")
	fmt.Scan(&coniugato)

	if coniugato {
		if reddito <= sogliaConiugato {
			tasse = float64(reddito) * aliquotaBassa
		} else {
			tasse = float64(sogliaConiugato)*aliquotaBassa + float64(reddito-sogliaConiugato)*aliquotaAlta
		}
		fmt.Printf("tasse per coniugato con reddito %d: %.2f\n", reddito, tasse)
	} else {
		if reddito <= sogliaNonConiugato {
			tasse = float64(reddito) * aliquotaBassa
		} else {
			tasse = float64(sogliaNonConiugato)*aliquotaBassa + float64(reddito-sogliaNonConiugato)*aliquotaAlta
		}
		fmt.Printf("tasse per non coniugato con reddito %d: %.2f\n", reddito, tasse)
	}
}


//Serve un solo if/else come minimo per distinguere se l'utente è coniugato o non coniugato, 
//poiché le soglie di reddito cambiano in base a questo stato.

//L'unico if necessario deve distinguere tra il caso "coniugato" e "non coniugato"