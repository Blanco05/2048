package main

import "fmt"

func main() {
    var x int
    var quad int
    fmt.Println("Immetti un valore di cui vuoi il quadrato")
    fmt.Scan(&x)
    quad = x * x
    fmt.Println("Il quadrato di", x, "è:", quad)
}
