package main

import (
    "fmt"
    "time"
    "os"
    "os/exec"
)

func main() {
    ticker := time.NewTicker(1 * time.Second)
    i := 0
    for {
        t := <-ticker.C
        fmt.Println(" ", t)
        time.Sleep(1 * time.Second)
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
        i++
        if i == 10{
            fmt.Println("Grazie per aver usato il programma!!")
            os.Exit(0)
        }
    }
}

