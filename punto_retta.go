package main

import (
    "fmt"
    "math"
)

func main() {
    const EPSILON = 1e-6
    var x1, y1, m, q float64
    fmt.Print("Inserisci le coordinate del punto P (x1, y1): ")
    fmt.Scan(&x1, &y1)
    fmt.Print("Inserisci i coefficienti della retta (m, q) per y = mx + q: ")
    fmt.Scan(&m, &q)

    y_retta := m*x1 + q

    if math.Abs(y1-y_retta) <= EPSILON {
        fmt.Println("sulla retta")
    } else if y1 > y_retta {
        fmt.Println("sopra")
    } else {
        fmt.Println("sotto")
    }
}
