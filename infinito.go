package main

import "fmt"
import "time"


func main(){
	var x float64
	x = 1.00001
	for {
		x += x
		fmt.Println(x)
		time.Sleep(100 * time.Millisecond)
	}
}