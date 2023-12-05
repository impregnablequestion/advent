package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "example.com/day_04/cards"
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
        card := scanner.Text()
        result, err := cards.ScoreCard(card)
        if err != nil {
            log.Fatalf("Fatal error with cards function: %v", err)
        }

        sum += result
        lines_read++
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Sum: %v , Lines read: %v", sum, lines_read)
}
