package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "example.com/day_03/parts"
)

func main() {

    f, err := os.Open("./input")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close();

    scanner := bufio.NewScanner(f)

    var schematic []string

    for scanner.Scan() {
        line := scanner.Text()
        schematic = append(schematic, line)

        if err != nil {
            log.Fatal(err)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    result, err := parts.CalculateTotal(schematic)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Result is %v", result)
}
