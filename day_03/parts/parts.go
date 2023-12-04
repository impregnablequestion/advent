package parts

import (
    "fmt"
	"log"
	"regexp"
	"strconv"
)

type number struct {
    content     string 
    row         int
    startCol    int
    endCol      int
}

const FullStop rune = '.'

func CalculateTotal(schematic []string) (uint, error) {

    var numbers []number
    var schemeMap [][]rune

    for i, line := range schematic {

        var currentNum []rune 
        var currentStart int
        var numbering bool

        schemeMap = append(schemeMap, []rune {})

        for j, char := range line {
            isNumber, err := isNumber(char)
            if err != nil {
                return 0, err
            }
            
            if isNumber {
                if !numbering {
                    currentStart = j - 1;
                    if j == 0 {
                        currentStart = 0
                    }
                    numbering = true
                }

                currentNum = append(currentNum, char)
                schemeMap[i] = append(schemeMap[i], 'n') 

                if j == len(line) - 1 {
                    numbering = false
                    numbers = append(numbers, number {
                        content:    string(currentNum),
                        row:        i,
                        startCol:   currentStart,
                        endCol:     j,
                    })
                    currentNum = []rune {}
                }
                continue
            }

            if numbering {
                numbering = false
                numbers = append(numbers, number {
                    content:    string(currentNum),
                    row:        i,
                    startCol:   currentStart,
                    endCol:     j,
                })

                currentNum = []rune {}
            }

            if char == FullStop {
                schemeMap[i] = append(schemeMap[i], char) 
                continue
            }

            schemeMap[i] = append(schemeMap[i], 's')
        }
    }

    // Now remove all the numbers without adjacent symbols from the calc

    validNumbers := []int {}

    for _, val := range schemeMap {
        for _, val := range val {
            fmt.Printf("%c", val)
        }
        fmt.Print("\n")
    }


    for _, num := range numbers {

        log.Printf("Number: %v , Start: %v , End: %v", num.content, num.startCol, num.endCol)

        // before and after
        if schemeMap[num.row][num.startCol] == 's' || schemeMap[num.row][num.endCol] == 's' {
            number, err := strconv.Atoi(num.content)

            if err != nil {
                log.Fatalf("Couldn't parse the number")
            }

            validNumbers = append(validNumbers, number)
            continue
        }

        // row above and below
        for i := num.startCol; i <= num.endCol; i++ {

            var rowAbove rune;
            var rowBelow rune;

            if (num.row > 0) {
                rowAbove = schemeMap[num.row - 1][i]
            }

            if (num.row < len(schemeMap) - 1) {
                rowBelow = schemeMap[num.row + 1][i]
            }

            if rowAbove == 's' || rowBelow == 's' {
                number, err := strconv.Atoi(num.content)

                if err != nil {
                    log.Fatalf("Couldn't parse the number")
                }

                validNumbers = append(validNumbers, number)
                break
            }
        }
    }

    var sum uint = 0

    for _, n := range validNumbers {
        sum += uint(n)
    }

    return sum, nil
}

func isNumber(char rune) (bool, error) {
    var bytes []byte
    bytes = append(bytes, byte(char))

    return regexp.Match("[0-9]", bytes)
}



