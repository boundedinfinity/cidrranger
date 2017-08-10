package main

import (
    "os"
    "log"
    "github.com/boundedinfinity/cidrranger/calc"
)

func main() {
    logger := log.New(os.Stdout, "", log.Lshortfile)
    service, err := calc.NewCalculatorService(
        calc.InputPath("./subnet-list-1.txt"),
        calc.Logger(logger),
        calc.Debug(true),
    )

    if err != nil {
        logger.Fatal(err)
        return
    }

    if highlow, err := service.FullRange(); err != nil {
        logger.Fatal(err)
    } else {
        logger.Printf("l[%s] - h[%s]", highlow[0], highlow[1])
    }
}


