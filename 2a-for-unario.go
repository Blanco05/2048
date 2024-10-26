package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
    /*
     scrivere un programma che date due costanti TARGET = 20 e MAX = 10 generi numeri casuali da 1 a MAX, lo stampi e poi lo sommi ai numeri già generati, se questa somma è maggiore di TARGET allora uscire dal ciclo e stampare la variabile somma 
    */
    const TARGET = 20
    const MAX = 10
    //rand.Seed(time.Now().Unix())
    var n int

    t := 0
    for t < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        t += n
    }
    fmt.Print("\n",t,"\n")
}
