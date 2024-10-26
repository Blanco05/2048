package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"time"
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

func creascacchiera() [][]int {
	matrix := make([][]int, 8)
	for i := range matrix {
		matrix[i] = make([]int, 8)
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i == 1 || i == 6 {
				matrix[i][j] = 2
			} else if i == 0 && (j == 0 || j == 7) || i == 7 && (j == 0 || j == 7) {
				matrix[i][j] = 3
			} else if i == 0 && (j == 1 || j == 6) || i == 7 && (j == 1 || j == 6) {
				matrix[i][j] = 4
			} else if i == 0 && (j == 2 || j == 5) || i == 7 && (j == 2 || j == 5) {
				matrix[i][j] = 5
			} else if i == 0 && j == 3 || i == 7 && j == 3 {
				matrix[i][j] = 6
			} else if i == 0 && j == 4 || i == 7 && j == 4 {
				matrix[i][j] = 7
			} else {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func printscacchiera(mat [][]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if mat[i][j] == 2 {
				fmt.Print("P   ")
			} else if mat[i][j] == 3 {
				fmt.Print("T   ")
			} else if mat[i][j] == 4 {
				fmt.Print("C   ")
			} else if mat[i][j] == 5 {
				fmt.Print("A   ")
			} else if mat[i][j] == 6 {
				fmt.Print("R   ")
			} else if mat[i][j] == 7 {
				fmt.Print("Re  ")
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println()
		fmt.Println()
	}
}

func movimento(mossa string, mat [][]int) {
	re := regexp.MustCompile(`([TCARD]?)(x?)([a-h][1-8])`)
	matches := re.FindStringSubmatch(mossa)

	if len(matches) == 0 {
		fmt.Println("Mossa non valida")
		return
	}

	pezzo := matches[1]       
	cattura := matches[2] == "x"
	destinazione := matches[3]

	col := int(destinazione[0] - 'a')
	riga := int('8' - destinazione[1])

	fmt.Println("Pezzo:", pezzo)
	fmt.Println("Cattura:", cattura)
	fmt.Println("Destinazione:", destinazione, "-> [", riga, ",", col, "]")

	if pezzo == "P" {
		P(mat, riga, col)
	} else if pezzo == "C" {
		C(mat, riga, col)
	}
}

func P(mat [][]int, riga, col int) {
	mat[6][4] = 0 
	mat[riga][col] = 2 
}

func C(mat [][]int, riga, col int) {
	mat[7][1] = 0
	mat[riga][col] = 4
}


func T(mat [][]int, riga, col int) {}
func A(mat [][]int, riga, col int) {}
func R(mat [][]int, riga, col int) {}
func Re(mat [][]int, riga, col int) {}

func main() {
	var mat [][]int
	var mossa string
	var flag = 0
	if flag == 0 {
		mat = creascacchiera()
		flag = 1
	}
	printscacchiera(mat)

	fmt.Println("Immetti la tua mossa: ")
	fmt.Scan(&mossa)

	movimento(mossa, mat)

	time.Sleep(4 * time.Second)
	clearTerminal()
	printscacchiera(mat)
}
