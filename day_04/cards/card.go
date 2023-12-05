package cards

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func ScoreCard(card string) (uint, error) {

    var winningNums uint = 0

    content := strings.Split(card, ": ")[1]
    
    log.Println(card)

    split := strings.Split(content, "|")
    winning, err := convertStringToNumbers(split[0])
    if err != nil {
        return 0, err
    }
    yours, err := convertStringToNumbers(split[1])
    if err != nil {
        return 0, err
    }

    for _, num := range yours {
        if isWin(winning, num) {
            winningNums++
            log.Printf("Win: %v", num)
        }
    }
    log.Print("\n")

    if winningNums == 0 {
        return 0, nil
    }

    return uint(math.Exp2(float64(winningNums - 1))), nil
}

func convertStringToNumbers(s string) ([]int, error) {
    numbers := []int {}
    split := strings.Fields(s);
    for _, val := range split {
        num, err := strconv.Atoi(val)
        if err != nil {
            return numbers, nil
        }
        numbers = append(numbers, num)
    }

    return numbers, nil
}

func isWin(winning []int, number int) bool {
    for i := range winning {
        if winning[i] == number {
            return true
        }
    }
    return false 
}
