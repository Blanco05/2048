package main

import "fmt"

func main(){
	var x, y float64
	fmt.Scan(&x, &y)
	if x > 0 && y > 0{
		fmt.Println("I quadrante")
	} 
	if x > 0 && y < 0 {
		fmt.Println("IV quadrante")
	} 
	if x < 0 && y < 0 {
		fmt.Println("III quadrante")
	}
	if x < 0 && y > 0 {
		fmt.Println("II quadrante")
	}
	if x == 0 && y < 0 {
		fmt.Println("IV quadrante")
	}
	if x == 0 && y > 0 {
		fmt.Println("I quadrante")
	}
	if x > 0 && y == 0 {
		fmt.Println("I quadrante")
	}
	if x < 0 && y == 0 {
		fmt.Println("II quadrante")
	}
	if x == 0 && y == 0{
		fmt.Println("I quadrante")
	}
}