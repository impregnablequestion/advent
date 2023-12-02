package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"example.com/day_01/extract"
)

func main() {

    f, err := os.Open("./calibration_file")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close();

    scanner := bufio.NewScanner(f)

    sum := 0
    lines_read := 0

    for scanner.Scan() {
        num := extract.Extract(scanner.Text())
        sum += num
        lines_read++
        fmt.Printf("Input: %v , Output: %v , Sum: %v\n", scanner.Text(), num, sum)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Sum is %v and a total of %v lines were read", sum, lines_read)
}
