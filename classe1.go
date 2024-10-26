package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number between 1 and 10:", randomInt(1, 10))
}
