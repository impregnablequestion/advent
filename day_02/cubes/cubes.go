package cubes

import (
    "strings"
    "log"
    "strconv"
)

const RedMax uint = 12;
const GreenMax uint = 13;
const BlueMax uint = 14;

func Cube(game string) (uint, error) {
    // Get the ID and the content 
    split := strings.Split(game, ": ")
    strId := strings.Replace(split[0], "Game ", "", 1)
    id, err := strconv.Atoi(strId)
    if err != nil {
        log.Fatalf("Error parsing id: %v", err)
        return 0, err
    }
    game = split[1]

    // Get the sessions

    sessions := strings.Split(game, "; ");
    for _, session := range sessions {
        log.Printf(session)
        cubes := strings.Split(session, ", ")

        for _, cube := range cubes {
            temp := strings.Fields(cube)

            i, err := strconv.Atoi(temp[0])
            if err != nil {
                log.Fatalf("Error parsing number: %v", err)
                return uint(id), err
            }

            number := uint(i)

            switch color := temp[1]; color {
            case "red":
                if (number > RedMax) {
                    return 0, err
                }
            case "blue":
                if (number > BlueMax) {
                    return 0, err
                }
            case "green":
                if (number > GreenMax) {
                    return 0, err
                }
            }
        }
    }

    return uint(id), err
}
