package main

import (
    "fmt"
    "math"
)

func main() {
    var x int
    var y float64

    // 1. int uguale a 10
    fmt.Print("int uguale a 10: ")
    fmt.Scan(&x)
    fmt.Println(x == 10)

    // 2. int diverso da 10
    fmt.Print("int diverso da 10: ")
    fmt.Scan(&x)
    fmt.Println(x != 10)

    // 3. int diverso da 10 e da 20
    fmt.Print("int diverso da 10 e da 20: ")
    fmt.Scan(&x)
    fmt.Println(x != 10 && x != 20)

    // 4. int diverso da 10 o da 20
    fmt.Print("int diverso da 10 o da 20: ")
    fmt.Scan(&x)
    fmt.Println(x != 10 || x != 20)

    // 5. int maggiore o uguale a 10
    fmt.Print("int maggiore o uguale a 10: ")
    fmt.Scan(&x)
    fmt.Println(x >= 10)

    // 6. int compreso tra 10 e 20, nell'intervallo [10,20]
    fmt.Print("int compreso tra 10 e 20 [10,20]: ")
    fmt.Scan(&x)
    fmt.Println(x >= 10 && x <= 20)

    // 7. int compreso tra 10 e 20, nell'intervallo (10,20)
    fmt.Print("int compreso tra 10 e 20 (10,20): ")
    fmt.Scan(&x)
    fmt.Println(x > 10 && x < 20)

    // 8. int compreso tra 10 e 20, nell'intervallo [10,20)
    fmt.Print("int compreso tra 10 e 20 [10,20): ")
    fmt.Scan(&x)
    fmt.Println(x >= 10 && x < 20)

    // 9. int esterno all'intervallo [10,20]
    fmt.Print("int esterno all'intervallo [10,20]: ")
    fmt.Scan(&x)
    fmt.Println(x < 10 || x > 20)

    // 10. int tra 10 e 20 [10,20] o tra 30 e 40 [30,40]
    fmt.Print("int tra 10 e 20 [10,20] o tra 30 e 40 [30,40]: ")
    fmt.Scan(&x)
    fmt.Println((x >= 10 && x <= 20) || (x >= 30 && x <= 40))

    // 11. int multiplo di 4 ma non di 100
    fmt.Print("int multiplo di 4 ma non di 100: ")
    fmt.Scan(&x)
    fmt.Println(x%4 == 0 && x%100 != 0)

    // 12. int dispari e compreso tra 0 e 100 [0,100]
    fmt.Print("int dispari e compreso tra 0 e 100 [0,100]: ")
    fmt.Scan(&x)
    fmt.Println(x%2 != 0 && x >= 0 && x <= 100)

    // 13. float64 vicino a 10.0 (non più lontano di 10^-6)
    fmt.Print("float64 vicino a 10.0 (non più lontano di 1e-6): ")
    fmt.Scan(&y)
    fmt.Println(math.Abs(y-10.0) <= 1e-6)
}
