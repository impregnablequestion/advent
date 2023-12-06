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

    // part 1

    cardList := []string {}

    for scanner.Scan() {
        card := scanner.Text()
        cardList = append(cardList, card)

        result, err := cards.ScoreCard(card)
        if err != nil {
            log.Fatalf("Fatal error with cards function: %v", err)
        }

        sum += result
        lines_read++
    }

    fmt.Printf("Sum: %v , Lines read: %v", sum, lines_read)
    
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // part 2 

    result, err := cards.ProcessCards(cardList)

    fmt.Printf("Cards counted: %v", result)



}
