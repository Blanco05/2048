package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
	for {
	    var rows, cols int
	    fmt.Println("\tMenù:\n1. Svastica (Tipo 1)\n2. Svastica (Tipo 2)\n3. Circonferenza\n4. Ellisse\n5. Quadrato\n6. Cubo\n7. Triangolo\n8. Clear")
	    fmt.Println("Immetti che tipo di Figura vuoi (1, 2, 3, 4, 5, 6, 7, 8): ")
	    var flag int
	    fmt.Scan(&flag)

	    if flag == 1 {
	        fmt.Println("Quanto deve essere grande la matrice (righe): ")
	        fmt.Scan(&rows)
	        fmt.Println("Quanto deve essere grande la matrice (colonne): ")
	        fmt.Scan(&cols)
	        mat := make([][]int, rows)
	        for i := range mat {
	            mat[i] = make([]int, cols)
	        }
	        for r := 0; r < rows; r++ {
	            for c := 0; c < cols; c++ {
	                if r < (rows-1)/2 && c == 0 || r == (rows-1)/2 || c == (cols-1)/2 || c > (cols-1)/2 && r == 0 || r == rows-1 && c < (cols-1)/2 || c == cols-1 && r > (rows-1)/2 {
	                    fmt.Print("x ")
	                } else {
	                    fmt.Print("  ")
	                }
	            }
	            fmt.Println()
	        }
	    }

	    if flag == 2 {
	        fmt.Println("Quanto deve essere grande la matrice (righe): ")
	        fmt.Scan(&rows)
	        fmt.Println("Quanto deve essere grande la matrice (colonne): ")
	        fmt.Scan(&cols)
	        mat := make([][]int, rows)
	        for i := range mat {
	            mat[i] = make([]int, cols)
	        }
	        for r := 0; r < rows; r++ {
	            for c := 0; c < cols; c++ {
	                if r > (rows-1)/2 && c == 0 || r == (rows-1)/2 || c == (cols-1)/2 || c < (cols-1)/2 && r == 0 || r == rows-1 && c > (cols-1)/2 || c == cols-1 && r < (rows-1)/2 {
	                    fmt.Print("x ")
	                } else {
	                    fmt.Print("  ")
	                }
	            }
	            fmt.Println()
	        }
	    }

	    if flag == 3 {
	        fmt.Println("Immetti il raggio della circonferenza: ")
	        var raggio int
	        fmt.Scan(&raggio)
	        mat := make([][]int, 2*raggio+1)
	        for i := range mat {
	            mat[i] = make([]int, 2*raggio+1)
	        }
	        for r := 0; r < 2*raggio+1; r++ {
	            for c := 0; c < 2*raggio+1; c++ {
	                if (r-raggio)*(r-raggio)+(c-raggio)*(c-raggio) <= raggio*raggio {
	                    fmt.Print("x ")
	                } else {
	                    fmt.Print("  ")
	                }
	            }
	            fmt.Println()
	        }
	    }

	    if flag == 4 {
	        fmt.Println("Quanto deve essere grande la matrice (righe): ")
	        fmt.Scan(&rows)
	        fmt.Println("Quanto deve essere grande la matrice (colonne): ")
	        fmt.Scan(&cols)
	        a := rows / 2
	        b := cols / 2
	        for r := 0; r < rows; r++ {
	            for c := 0; c < cols; c++ {
	                if (r-a)*(r-a)*b*b + (c-b)*(c-b)*a*a == a*a*b*b {
	                    fmt.Print("x ")
	                } else {
	                    fmt.Print("  ")
	                }
	            }
	            fmt.Println()
	        }
	    }

	    if flag == 5 {
	    	fmt.Println("Quanto deve essere grande il quadrato: ")
	        fmt.Scan(&rows)
	        for r := 0; r < rows; r++ {
	            for c := 0; c < rows; c++ {
	                if r == rows-1 || c == rows-1 || c == 0 || r == 0{
	                    fmt.Print("x ")
	                } else {
	                    fmt.Print("  ")
	                }
	            }
	            fmt.Println()
	        }
	    }

	    if flag == 6 {
	    	fmt.Println("Quanto deve essere grande il quadrato: ")
	        fmt.Scan(&rows)
	        h := rows/2
	        sum := h + rows
	        i := sum-1
	        for r := 0; r < sum; r++ {
	            for c := 0; c < sum; c++ {
	                if r == sum-1 && c <h*2 || c == rows-1 && r > h || c == 0 && r > h || r == h && c <h*2 || r==0 && c>h || c == sum-1 && r<= (h*2)-1 || c == i && c > (h*2)-1 || c > (h*2)-1 && c == i + (h*2)-1|| r < (h*2) && c == i - (h*2)+1{
	                    fmt.Print("x ")
	                } else {
	                    fmt.Print("  ")
	                }
	            }   
	            fmt.Println()
	            i--
	        }
	    }
	    if flag == 7 {
	    	fmt.Println("Quanto deve essere grande la matrice (righe): ")
	        fmt.Scan(&rows)
	        i := 0
	        for r := 0; r < rows; r++ {
	            for c := 0; c < rows; c++ {
	                if c == i || c == 0 || r == rows - 1{
	                    fmt.Print("x ")
	                } else {
	                    fmt.Print("  ")
	                }
	            }
	            i++
	            fmt.Println()
	        }
	    }
	    if flag == 8{
	    	cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
	    }
		fmt.Println("Se vuoi uscire dal programma premi 1, se vuoi rimanere premi 0: ")
		tot := 0
		fmt.Scan(&tot)
		if tot == 1{
			fmt.Println("Grazie per aver usato il programma!!")
			os.Exit(0)
		}
	}
}