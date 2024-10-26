package main

import "fmt"

func main() {
    var n int
    fmt.Print("Inserisci un numero positivo: ")
    fmt.Scan(&n)

    if n <= 0 {
        fmt.Println("Il numero deve essere positivo.")
        return
    }

    for i := 1; i*i <= n; i++ {
        quadrato := i * i
        fmt.Println(quadrato)
    }
}
