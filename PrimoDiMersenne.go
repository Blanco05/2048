package main

import (
	"fmt"
	"math"
)

// Funzione per verificare se un numero intero è primo
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Funzione per calcolare il numero di Mersenne dato un esponente n
func calcola(n int) int {
	return int(math.Pow(2, float64(n))) - 1
}

func main() {
	for n := 2; ; n++ {
		if IsPrime(n) {  
			numero := calcola(n) 
			if IsPrime(numero) { 
				fmt.Println("\033[32m", numero, "\033[0m")
			}
		}
	}
}
