package main

import "fmt"

func main() {
	hand := []string{
		"         ____  ",
		"        |    | ",
		"        |    | ",
		"        |    | ",
		"        | ❤  | ",
		"        |    | ",
		"    ____|    |_____",
		"   |               |",
		"   |               |",
		"   |_______________|",
		"        |     |     ",
		"        |     |     ",
		"        |     |     ",
		"        |_____|     ",
	}

	for _, line := range hand {
		fmt.Println(line)
	}
}
