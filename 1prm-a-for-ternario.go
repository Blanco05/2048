package main
import "fmt"
func main() {
    /*
     Far Immettere all'utente un numero, e stamoare tutti i numeri parti sotto n compreso lo zero
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
