package cubes

import (
    "strings"
    "log"
    "strconv"
)

func Cube(game string) (uint, error) {
    // Remove the ID and the content 
    split := strings.Split(game, ": ")
    game = split[1]

    var redMin uint = 0;
    var blueMin uint = 0;
    var greenMin uint = 0;

    // Get the sessions

    sessions := strings.Split(game, "; ");
    for _, session := range sessions {

        cubes := strings.Split(session, ", ")

        for _, cube := range cubes {
            temp := strings.Fields(cube)

            i, err := strconv.Atoi(temp[0])
            if err != nil {
                log.Fatalf("Error parsing number: %v", err)
                return 0, err
            }

            number := uint(i)

            switch color := temp[1]; color {
            case "red":
                if (number > redMin) {
                    redMin = number
                }
            case "blue":
                if (number > blueMin) {
                    blueMin = number 
                }
            case "green":
                if (number > greenMin) {
                    greenMin = number
                }
            }
        }
    }

    power := redMin * blueMin * greenMin 
    return power, nil 
}
