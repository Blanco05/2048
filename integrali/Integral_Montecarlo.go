package main

import "fmt"
import "math/rand"
import "math"

/*
 * il metodo montecralo funziona creando un numero elevato di punti casuali
 * in un rettangolo che comprende i bordi dell'integrale definito
 * e poi si vede quanti di questi punti sono sopra e sotto il grafico e 
 * poi si calcola in percentuale l'area del rettangolo in base ai punti
 * presi
 */

func IntegraMonteCarlo(f func(float64) float64, a float64, b float64, n int, M float64) float64 {
	sotto := 0
	for i:=0; i<n; i++{
		x := a + rand.Float64() * (b-a)
		y := rand.Float64() * M
		if y < f(x){
			sotto++
		}
	}
	areaPos := M * (b-a) * float64(sotto)/float64(n)
	sopra := 0
	for i:=0; i<n; i++{
		x := a + rand.Float64() * (b-a)
		y := -rand.Float64() * M
		if y > f(x){
			sopra++
		}
	}
	areaNeg := M * (b-a) * float64(sopra)/float64(n)
	return areaPos - areaNeg
}

func main(){
	actual := IntegraMonteCarlo(math.Sin, 3, 7, 100000, 2)
	expected := -math.Cos(7)-(-math.Cos(3))
	fmt.Println(actual, expected)
}