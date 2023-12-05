package parts

import (
    "fmt"
    "log"
    "regexp"
    "strconv"
)

type xy struct {
    x       int
    y       int
}

type number struct {
    content         string 
    row             int
    startCol        int
    endCol          int
    contentSquares  []xy
}

type gear struct {
    startRow        int
    endRow          int
    startCol        int
    endCol          int
}

const FullStop rune = '.'

func CalculateTotalGearRatios(schematic []string) (uint, error) {
    numbers, schemeMap, err := getNumbersAndSchemeMap(schematic)
    if err != nil {
        return 0, err
    }

    printSchemeMap(schemeMap)

    gears := getGears(schematic)
    sum := 0;

    for _, gear := range gears {
        gearNumbers := []int {}

        for i := gear.startRow; i <= gear.endRow; i++ {
            numbering := false

            for j := gear.startCol; j <= gear.endCol; j++ {

                fmt.Printf("%c", schemeMap[i][j])

                if schemeMap[i][j] == 'n' && !numbering {
                    numbering = true

                    for _, num := range numbers {
                        for _, xy := range num.contentSquares {
                            if xy.x == j && xy.y == i  {

                                number, err := strconv.Atoi(num.content)
                                if err != nil {
                                    log.Fatalf("Couldn't parse the number")
                                }
                                
                                gearNumbers = append(gearNumbers, number)
                            }
                        }
                    }
                } else {
                    numbering = false
                }
            }

            fmt.Print("\n")
        }

        log.Printf("Gear numbers: %v", gearNumbers)

        if len(gearNumbers) == 2 {
            sum += (gearNumbers[0] * gearNumbers[1])
        }
    }


    return uint(sum), nil
}

func CalculateTotal(schematic []string) (uint, error) {

    numbers, schemeMap, err := getNumbersAndSchemeMap(schematic)
    if err != nil {
        return 0, err
    }

    // Now remove all the numbers without adjacent symbols from the calc

    validNumbers := []int {}

    for _, num := range numbers {

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
    for _, n := range validNumbers { sum += uint(n) }
    return sum, nil
}

func isNumber(char rune) (bool, error) {
    var bytes []byte
    bytes = append(bytes, byte(char))

    return regexp.Match("[0-9]", bytes)
}

func getNumbersAndSchemeMap(schematic []string) ([]number, [][]rune, error) {

    numbers := []number {}
    schemeMap := [][]rune {}

    for i, line := range schematic {

        var currentNum []rune 
        var currentXY []xy
        var currentStart int
        var numbering bool

        schemeMap = append(schemeMap, []rune {})

        for j, char := range line {
            isNumber, err := isNumber(char)
            if err != nil {
                return []number {}, [][]rune {}, err
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
                currentXY = append(currentXY, xy { x: j, y: i })
                schemeMap[i] = append(schemeMap[i], 'n') 

                if j == len(line) - 1 {
                    numbering = false
                    numbers = append(numbers, number {
                        content:        string(currentNum),
                        row:            i,
                        startCol:       currentStart,
                        endCol:         j,
                        contentSquares: currentXY,
                    })
                    currentNum = []rune {}
                    currentXY = []xy {}
                }
                continue
            }

            if numbering {
                numbering = false
                numbers = append(numbers, number {
                    content:        string(currentNum),
                    row:            i,
                    startCol:       currentStart,
                    endCol:         j,
                    contentSquares: currentXY,
                })

                currentNum = []rune {}
                currentXY = []xy {}
            }

            if char == FullStop {
                schemeMap[i] = append(schemeMap[i], char) 
                continue
            }

            schemeMap[i] = append(schemeMap[i], 's')
        }
    }

    return numbers, schemeMap, nil   
}

func getGears(schematic []string) []gear {
    gears := []gear {}
    for i, line := range schematic {
        for j, char := range line {
            if char == '*' {

                startRow := i - 1
                endRow := i + 1
                startCol := j - 1
                endCol := j + 1

                if startRow < 0 {
                    startRow = 0
                }

                if startCol < 0 {
                    startCol = 0
                }

                if endRow == len(schematic) {
                    endRow--
                }

                if endCol == len(schematic) {
                    endCol--
                }

                gears = append(gears, gear {
                    startRow:   startRow,
                    endRow:     endRow,
                    startCol:   startCol,
                    endCol:     endCol,
                })
            }
        }
    }

    return gears
}

func printSchemeMap(schemeMap [][]rune) {

    for _, val := range schemeMap {
        for _, val := range val {
            fmt.Printf("%c", val)
        }
        fmt.Print("\n")
    }
}
