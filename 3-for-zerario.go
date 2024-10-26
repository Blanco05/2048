package main
import (
    "fmt"
    "math/rand"
    "time"
)
func main() {
    /*
     Scivere un programma che dato un K=20, generi numeri casuali da 0 a K, li stampi, che implementi una variabile e che se il numero è uguale a zero esca dal ciclo
    */
    //rand.Seed(time.Now().Unix())
    const K = 20
    var n int

    c := 0
    for {
        n = rand.Intn(K+1)
        fmt.Print(n, " ")
        if n == 0 {
            break
        }
        c++
    }
    fmt.Println(c)
}
