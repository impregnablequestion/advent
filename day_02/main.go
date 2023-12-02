package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "example.com/day_02/cubes"
)

func main() {

    f, err := os.Open("./input")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close();

    scanner := bufio.NewScanner(f)

    sum := 0
    lines_read := 0

    for scanner.Scan() {
        game := scanner.Text()
        validGame := cubes.Cube(scanner.Text())

        if validGame {
            sum++
        }

        lines_read++
        fmt.Printf("Input: %v , Result: %v, Sum: %v\n", game, validGame, sum)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Sum: %v , Lines read: %v", sum, lines_read)
}
