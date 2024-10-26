package main

import (
	"fmt"
	"runtime"
	"os/exec"
	"os"
	"time"
	"math/rand"
	"log"
	"github.com/nsf/termbox-go"
)

const width, height = 20, 10

func tastiera() rune {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			return ev.Ch
		case termbox.EventError:
			log.Fatal(ev.Err)
		}
	}
}

func randy() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := height - 2
	return rand.Intn(max-min+1) + min
}

func randx() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := width - 2
	return rand.Intn(max-min+1) + min
}

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

func stampa(mat [][]int, melax, melay int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if mat[i][j] == 9 {
				fmt.Print("\033[32mS \033[0m")
			} else if i == melay && j == melax {
				fmt.Print("\033[31mo \033[0m")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func muoviSerpentey(mat [][]int, x, y int, direction int) (int, int) {
	mat[y][x] = 1
	y += direction
	if y >= height-1 {
		y = 1
	} else if y < 1 {
		y = height - 2
	}
	mat[y][x] = 9
	return x, y
}

func muoviSerpentex(mat [][]int, x, y int, direction int) (int, int) {
	mat[y][x] = 1
	x += direction
	if x >= width-1 {
		x = 1
	} else if x < 1 {
		x = width - 2
	}
	mat[y][x] = 9
	return x, y
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	mat := make([][]int, height)
	for i := range mat {
		mat[i] = make([]int, width)
		for j := 0; j < width; j++ {
			if i == 0 || j == 0 || j == width-1 || i == height-1 {
				mat[i][j] = 0
			} else {
				mat[i][j] = 1
			}
		}
	}

	x, y := 19, 1
	mat[y][x] = 9

	melax, melay := 19, 2

	directionX, directionY := 0, 1

	for {
		clearTerminal()
		stampa(mat, melax, melay)
		tasto := tastiera()

		switch tasto {
		case 'w':
			directionX, directionY = 0, -1
		case 's':
			directionX, directionY = 0, 1
		case 'a':
			directionX, directionY = -1, 0
		case 'd':
			directionX, directionY = 1, 0
		}

		if directionX != 0 {
			x, y = muoviSerpentex(mat, x, y, directionX)
		} else if directionY != 0 {
			x, y = muoviSerpentey(mat, x, y, directionY)
		}

		time.Sleep(200 * time.Millisecond)
	}
}
