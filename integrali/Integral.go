package main

import "fmt"
import "math"

func IntegraTrap(f func(float64) float64, a float64, b float64, n int) float64 {
	delta := (b-a) / float64(n)
	somma := 0.0
	for i := 0; i<n ; i++{
		start := a + float64(i) * delta
		end := a + float64(i+1) * delta
		areaTrapezio := (f(start) + f(end)) * delta / 2.0
		somma += areaTrapezio
		fmt.Println(somma)
	}
	return somma
}

func main(){
	actual := IntegraTrap(math.Sin, 3, 7, 1000000)
	expected := -math.Cos(7)-(-math.Cos(3))
	fmt.Println(actual, expected)
}