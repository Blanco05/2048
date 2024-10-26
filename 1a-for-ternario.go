package main
import "fmt"
func main() {
    /*
	 Richiedere un numero all'utente e scrivere tanti * quanti il numero dell'utente
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
