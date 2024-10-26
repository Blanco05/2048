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
    for i:=0; i<10; i++{
    	    fmt.Println(rand.Intn(max-min+1) + min)
    }
}	