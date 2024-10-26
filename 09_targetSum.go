package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
    const TARGET = 50
    const MAX = 10
    rand.Seed(time.Now().Unix())
    var n int

    t := 0
    for t < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        t += n
    }
    fmt.Print("\n",t,"\n")
}
