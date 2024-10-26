package main

import(
	"fmt"
)

func main(){
	var x = 11
	mat := make([][]int, 6)
	for i := range mat {
		mat[i] = make([]int, 11)
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 11; j++ {
			if j <= x/2 + i &&  j >= x/2 - i{
				fmt.Print("# ")
			} else if i == 5 {
				fmt.Print("# ")
			}else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}