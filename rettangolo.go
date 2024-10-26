package main

import "fmt"

func main() {
	var x1, y1, x2, y2, x3, y3 int
	// Richiedo gli estremi del rettangolo
	fmt.Println("P1:")
	fmt.Scan(&x1, &y1)
	fmt.Println("P2:")
	fmt.Scan(&x2, &y2)
	// Richiedo il punto esterno
	fmt.Println("P3:")
	fmt.Scan(&x3, &y3)

	// Calcolo i minimi e massimi delle coordinate
	minX := x1
	maxX := x2
	minY := y1
	maxY := y2

	if x1 > x2 {
		minX, maxX = x2, x1
	}
	if y1 > y2 {
		minY, maxY = y2, y1
	}

	// Verifico se P3 è esterno al rettangolo
	if x3 < minX || x3 > maxX || y3 < minY || y3 > maxY {
		fmt.Println("P3 esterno")
	} else {
		fmt.Println("P3 interno")
	}
}
