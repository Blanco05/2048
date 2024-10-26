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
    somma := 0
    for i:=0; i<10; i++{
    	n := rand.Intn(max-min+1) + min
    	fmt.Print(n, " ")
    	somma += n
    }
    fmt.Print("\nLa loro somma è: ", somma)
}