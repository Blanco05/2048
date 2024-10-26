package main
import "fmt"
func main() {
   // Questo programma legge numeri in coppie e combina le loro ultime 2 cifre
   // Continua finché non ottiene una combinazione uguale a 100
   
   var numeroCombinato, numeroCorrente, numeroPrecedente int
   fmt.Scan(&numeroCorrente)
   for numeroCombinato != 100 {
       numeroPrecedente = numeroCorrente
       fmt.Scan(&numeroCorrente)
       numeroCombinato = (numeroPrecedente%100)*100 + numeroCorrente%100
       fmt.Println(numeroCombinato)
   }
}