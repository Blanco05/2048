package main

import "fmt"

func main() {
    const EPSILON = 1e-5
    var x, y float64
    fmt.Scan(&x, &y)

    if x*x + y*y < EPSILON*EPSILON {
        fmt.Println("molto vicino all'origine")
    } else {
        fmt.Println("non vicino all'origine")
    }
}
