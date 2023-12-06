package cards

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type processedCard struct {
    number          int
    nextCards       []int
}

func ProcessCards(cards []string) (int, error) {
    processedCards := []processedCard {}
    for _, card := range cards {
        result, err := processCard(card)
        if err != nil {
            return 0, err
        }

        processedCards = append(processedCards, result)
    }

    sum := 0

    for _, card := range processedCards {
        countCard(card, processedCards, &sum)
    }

    return sum, nil
}

func countCard(cardToCount processedCard, cards []processedCard, sum *int) {
    *sum++
    for _, cardNo := range cardToCount.nextCards {
        nextCard, err := findCard(cardNo, cards)
        if err != nil {
            break
        } 

        countCard(nextCard, cards, sum)
    }
}

func findCard(cardNo int, cards []processedCard) (processedCard, error) {
    for i := range cards {
        if cards[i].number == cardNo {
            return cards[i], nil
        }
    }

    return processedCard{}, errors.New("No corresponding card")
}

func processCard(card string) (processedCard, error){
    copies := []int {}

    split := strings.Split(card, ": ")
    prefix, content := split[0], split[1]
    cardNo, err := strconv.Atoi(strings.Fields(prefix)[1])
    if err != nil {
        return processedCard{}, err
    }

    result, err := getWinningNums(content)
    if err != nil {
        return processedCard{}, err
    }

    nextCard := cardNo + 1

    for i := 0; i < result; i++ {
        copy := nextCard + i
        copies = append(copies, copy)
    }

    processedCard := processedCard {
        number:     cardNo,
        nextCards:  copies,
    }

    return processedCard, nil
}

func getWinningNums(content string) (int, error) {

    winningNums := 0;

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
        }
    }

    return winningNums, nil
}

func ScoreCard(card string) (uint, error) {

    content := strings.Split(card, ": ")[1]

    wins, err := getWinningNums(content)

    if err != nil {
        return 0, err
    }

    if wins == 0 {
        return 0, nil
    }

    return uint(math.Exp2(float64(wins - 1))), nil
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
