package main

import "fmt"

func main() {
   var a, b int
   fmt.Scan(&a, &b)
   for b != 0 {
       r := a % b 
       a = b     
       b = r    
   }
   fmt.Println(a)
}