package cards

import (
    "testing"
    "fmt"
)

type test struct {
    want        uint 
    content     string
}

func TestMain (t *testing.T) {
    cases := []test {
        {
            want:       8,
            content:    "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
        },
        {
            want:       2,
            content:    "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
        },
        {
            want:       2,
            content:    "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
        },
        {
            want:       1,
            content:    "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
        },
        {
            want:       0,
            content:    "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
        },
        {
            want:       0,
            content:    "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
        },
    }

    for i, val := range cases {
        title := fmt.Sprintf("Case: %v", i + 1)
        t.Run(title, func(t *testing.T) {
            result, err := ScoreCard(val.content) 
            if err != nil {
                t.Fatalf(err.Error())
            }
            
            if result != val.want {
                t.Errorf("Wanted: %v Got: %v", val.want, result)
            }
        })
    }
}
