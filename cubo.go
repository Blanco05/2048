package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

func mat1() {
	cols := 11
	mat1 := make([][]string, cols)
	for i := range mat1 {
		mat1[i] = make([]string, cols)
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < cols; j++ {
			if i == 3 || i == 7 || j == 3 || j == 7 {
				fmt.Print("  ")
			} else {
				mat1[i][j] = "x"
				fmt.Print(mat1[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func mat2() {
	cols := 11
	mat2 := make([][]string, cols)
	for i := range mat2 {
		mat2[i] = make([]string, cols)
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < cols; j++ {
			if i == 3 || i == 7 || j == 3 || j == 7 {
				fmt.Print("  ")
			} else {
				mat2[i][j] = "\033[34mx\033[0m"
				fmt.Print(mat2[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func mat3() {
	cols := 11
	mat3 := make([][]string, cols)
	for i := range mat3 {
		mat3[i] = make([]string, cols)
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < cols; j++ {
			if i == 3 || i == 7 || j == 3 || j == 7 {
				fmt.Print("  ")
			} else {
				mat3[i][j] = "\033[32mx\033[0m"
				fmt.Print(mat3[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func mat4() {
	cols := 11
	mat4 := make([][]string, cols)
	for i := range mat4 {
		mat4[i] = make([]string, cols)
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < cols; j++ {
			if i == 3 || i == 7 || j == 3 || j == 7 {
				fmt.Print("  ")
			} else {
				mat4[i][j] = "\033[33mx\033[0m"
				fmt.Print(mat4[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func mat5() {
	cols := 11
	mat5 := make([][]string, cols)
	for i := range mat5 {
		mat5[i] = make([]string, cols)
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < cols; j++ {
			if i == 3 || i == 7 || j == 3 || j == 7 {
				fmt.Print("  ")
			} else {
				mat5[i][j] = "\033[38;5;208mx\033[0m"
				fmt.Print(mat5[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func mat6() {
	cols := 11
	mat6 := make([][]string, cols)
	for i := range mat6 {
		mat6[i] = make([]string, cols)
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < cols; j++ {
			if i == 3 || i == 7 || j == 3 || j == 7 {
				fmt.Print("  ")
			} else {
				mat6[i][j] = "\033[31mx\033[0m"
				fmt.Print(mat6[i][j], " ")
			}
		}
		fmt.Println()
	}
}

type Cubo struct {
	currentFace int
	movimenti   map[int]map[string]int
}

func newCubo() *Cubo {
	return &Cubo{
		currentFace: 0,
		movimenti: map[int]map[string]int{
			0: {"w": 1, "s": 2, "a": 4, "d": 3},
			1: {"w": 5, "s": 0, "a": 4, "d": 3},
			2: {"w": 0, "s": 5, "a": 4, "d": 3},
			3: {"w": 1, "s": 2, "a": 0, "d": 5},
			4: {"w": 1, "s": 2, "a": 5, "d": 0},
			5: {"w": 1, "s": 2, "a": 4, "d": 3},
		},
	}
}

func (c *Cubo) visualizzaFaccia() {
	clearTerminal()
	switch c.currentFace {
	case 0:
		mat1()
	case 1:
		mat2()
	case 2:
		mat3()
	case 3:
		mat4()
	case 4:
		mat5()
	case 5:
		mat6()
	}
}

func (c *Cubo) sposta(n string) {
	if nextFace, ok := c.movimenti[c.currentFace][n]; ok {
		c.currentFace = nextFace
	} else {
		fmt.Println("Movimento non valido!")
	}
}

func main() {
	cubo := newCubo() 

	for {
		cubo.visualizzaFaccia()

		var n string
		fmt.Print("Inserisci un carattere tra A, W, S, D: ")
		fmt.Scan(&n)

		cubo.sposta(n)

		var c int
		fmt.Print("Vuoi terminare il codice? (1 per Sì, 0 per No): ")
		fmt.Scan(&c)
		if c == 1 {
			os.Exit(0)
		}
	}
}
