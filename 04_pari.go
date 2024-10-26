package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
 	min := 0
    max := 10
    count := 0
    for i:=0; i<10; i++{
    	n := rand.Intn(max-min+1) + min
    	fmt.Print(n, " ")
    	if n % 2 == 0{
    		count ++
    	}
    }
    fmt.Print("\ni numeri pari sono: ", count)
}	