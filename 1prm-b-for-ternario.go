package main
import "fmt"
func main() {
    /*
     chiedere all'utente un numero in ingresso e stampare in ordine descrescente i numeri sotto esso
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

