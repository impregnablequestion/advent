package extract

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getNumber(written string) string {
    switch written {
    case "one":
        return "1"
    case "eno":
        return "1"
    case "two":
        return "2"
    case "owt":
        return "2"
    case "three":
        return "3"
    case "eerht":
        return "3"
    case "four":
        return "4"
    case "ruof":
        return "4"
    case "five":
        return "5"
    case "evif":
        return "5"
    case "six":
        return "6"
    case "xis":
        return "6"
    case "seven":
        return "7"
    case "neves":
        return "7"
    case "eight":
        return "8"
    case "thgie":
        return "8"
    case "nine":
        return "9"
    case "enin":
        return "9"
    default:
        return "0"
    }
}

func replaceFirst(line string, matches []string) string {
    return strings.Replace(line, matches[0], getNumber(matches[0]), 1)
}

func replaceLast(line string, matches[]string) string {
    toReplace := matches[0]
    return strings.Replace(line, toReplace, getNumber(toReplace), 1)
}

func reverseString(str string) string{
   byte_str := []rune(str)
   for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
      byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
   }
   return string(byte_str)
}

func Extract(line string) int {
    // replace first and last instance of written digit with real digit
    writtenDigits := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")
    matches := writtenDigits.FindStringSubmatch(line);

    fmt.Printf("Initial: %v \n", line)
    if matches != nil {
        line = replaceFirst(line, matches)
        fmt.Printf("Replaced first occurence: %v \n", line)
        
    }

    writtenDigitsRev := regexp.MustCompile("eno|owt|eerht|ruof|evif|xis|neves|thgie|enin")
    reverseMatches := writtenDigitsRev.FindStringSubmatch(reverseString(line));
    if reverseMatches != nil {
        line = replaceLast(reverseString(line), reverseMatches)
        line = reverseString(line)
        fmt.Printf("Replaced last occurence: %v \n", line)
    }

    // extract number from formatted string
    regex := regexp.MustCompile("[0-9]")
    digits := regex.FindAllString(line, -1);
    if len(digits) == 0 {
        return 0
    }

    number := digits[0] + digits[len(digits) - 1]

    i, err := strconv.Atoi(number);
    if err != nil {
        return 0
    }
    fmt.Printf("Answer: %v \n\n", i)

    return i
}
