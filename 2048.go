package main

import (
	"fmt"
	"runtime"
	"os/exec"
	"os"
	"math/rand"
)

func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func randompos(mat [][]int) (int, int) {
	x, y := rand.Intn(4), rand.Intn(4)
	for mat[x][y] != 0 { 
		x, y = rand.Intn(4), rand.Intn(4)
	}
	return x, y
}

func a(mat [][]int) {
	clearTerminal()
	for i := 0; i < 4; i++ {
		for j := 1; j < 4; j++ {
			if mat[i][j] != 0 {
				for k := j; k > 0 && mat[i][k-1] == 0; k-- {
					mat[i][k-1] = mat[i][k]
					mat[i][k] = 0
				}
				if j > 0 && mat[i][j-1] == mat[i][j] {
					mat[i][j-1] += mat[i][j]
					mat[i][j] = 0
				}
			}
		}
	}
	printMatrix(mat)
}

func w(mat [][]int) {
	clearTerminal()
	for j := 0; j < 4; j++ {
		for i := 1; i < 4; i++ {
			if mat[i][j] != 0 {
				for k := i; k > 0 && mat[k-1][j] == 0; k-- {
					mat[k-1][j] = mat[k][j]
					mat[k][j] = 0
				}
				if i > 0 && mat[i-1][j] == mat[i][j] {
					mat[i-1][j] += mat[i][j]
					mat[i][j] = 0
				}
			}
		}
	}
	printMatrix(mat)
}

func s(mat [][]int) {
	clearTerminal()
	for j := 0; j < 4; j++ {
		for i := 2; i >= 0; i-- {
			if mat[i][j] != 0 {
				for k := i; k < 3 && mat[k+1][j] == 0; k++ {
					mat[k+1][j] = mat[k][j]
					mat[k][j] = 0
				}
				if i < 3 && mat[i+1][j] == mat[i][j] {
					mat[i+1][j] += mat[i][j]
					mat[i][j] = 0
				}
			}
		}
	}
	printMatrix(mat)
}

func d(mat [][]int) {
	clearTerminal()
	for i := 0; i < 4; i++ {
		for j := 2; j >= 0; j-- {
			if mat[i][j] != 0 {
				for k := j; k < 3 && mat[i][k+1] == 0; k++ {
					mat[i][k+1] = mat[i][k]
					mat[i][k] = 0
				}
				if j < 3 && mat[i][j+1] == mat[i][j] {
					mat[i][j+1] += mat[i][j]
					mat[i][j] = 0
				}
			}
		}
	}
	printMatrix(mat)
}

func printMatrix(mat [][]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if countDigitsUsingDivision(mat[i][j]) == 1{
				fmt.Print("\033[33m", mat[i][j], "\033[0m", "      ")
			}
			if countDigitsUsingDivision(mat[i][j]) == 2{
				fmt.Print("\033[31m", mat[i][j], "\033[0m", "     ")
			}
			if countDigitsUsingDivision(mat[i][j]) == 3{
				fmt.Print("\033[34m", mat[i][j], "\033[0m", "    ")
			}
		}
		fmt.Println("\n\n\n")
	}
}

func countDigitsUsingDivision(n int) int {
    if n == 0 {
        return 1
    }
    count := 0
    if n < 0 {
        n = -n 
    }
    for n > 0 {
        n /= 10
        count++
    }
    return count
}

func isFull(mat [][]int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if mat[i][j] == 0 {
				return false 
			}
		}
	}
	return true 
}



func main() {
	mat := make([][]int, 4)
	for i := range mat {
		mat[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			mat[i][j] = 0 
		}
	}

	
	x, y := randompos(mat)
	mat[x][y] = 2 
	z, q := randompos(mat)
	mat[z][q] = 4
	f, g := randompos(mat)
	mat[f][g] = 2

	fmt.Println("Ecco a te la tua matrice di 2048:\n")
	printMatrix(mat)

	for {
		fmt.Println("Immetti un carattere tra a, w, s, d: ")
		var comando rune 
		fmt.Scanf("%c\n", &comando) 

		switch comando {
		case 'a':
			a(mat)
		case 'w':
			w(mat)
		case 's':
			s(mat)
		case 'd':
			d(mat)
		default:
			fmt.Println("Comando non valido")
		}

		if isFull(mat) {
			fmt.Println("La matrice è piena! Non ci sono più mosse disponibili.")
			break
		}

		
		x, y = randompos(mat)
		mat[x][y] = 2 
	}
}
