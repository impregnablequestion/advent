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

    var sum uint = 0;
    var lines_read uint = 0

    for scanner.Scan() {
        game := scanner.Text()
        result, err := cubes.Cube(game)
        if err != nil {
            log.Fatal(err)
        }
        sum += result

        lines_read++
        fmt.Printf("Input: %v , Result: %v, Sum: %v\n", game, result, sum)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Sum: %v , Lines read: %v", sum, lines_read)
}
