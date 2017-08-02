package main

import (
    "os"
    "log"
    "github.com/boundedinfinity/cidrranger/calc"
)

func main() {
    file, err := os.Open("./subnet-list-1.txt")
    //file, err := os.Open("./subnet-list-2.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    if highlow, err := calc.FullRange(file); err != nil {
        log.Fatal(err)
    } else {
        log.Printf("l[%s] - h[%s]", highlow[0], highlow[1])
    }
}


