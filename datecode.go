//date due date, stabilisci se la prima precede la seconda

package main

import ("fmt")

func main(){
	day1, month1, year1 := 0,0,0
	fmt.Print("immetti giorni data 1: ")
	fmt.Scan(&day1)

	fmt.Print("immetti mesi data 1: ")
	fmt.Scan(&month1)

	fmt.Print("immetti anni data 1: ")
	fmt.Scan(&year1)
	fmt.Println()
	day2, month2, year2 := 0,0,0
	fmt.Print("immetti giorni data 2: ")
	fmt.Scan(&day2)

	fmt.Print("immetti mesi data 2: ")
	fmt.Scan(&month2)

	fmt.Print("immetti anni data 2: ")
	fmt.Scan(&year2)

	if year2 > year1{
		fmt.Print("la data 2 è maggiore")
	}
	if year2 == year1 {
		if month1 == month2 {
			if day1 == day2 {
				fmt.Print("le due date sono uguali")
			}else if day1 > day2 {
				fmt.Print("la data 1 è maggiore")
			} else if day1 < day2{
				fmt.Print("la data 2 è maggiore")
			}
		} else if month1 > month2 {
			fmt.Print("la data 1 è maggiore")
		} else if month1 < month2{
			fmt.Print("la data 2 è maggiore")
		}
	}
	if year2 < year1{
		fmt.Print("la data 1 è maggiore")
	}
}