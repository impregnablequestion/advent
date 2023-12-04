package parts

import (
    "testing"
)

type test struct {
    name        string
    schematic   []string
    want        uint
}

func TestMain(t *testing.T) {
    schematic := []string {
        "467..114..",
        "...*......",
        "..35..633.",
        "......#...",
        "617*......",
        ".....+.58.",
        "..592.....",
        "......755.",
        "...$.*....",
        ".664.598..",
    }

    var expected uint = 4361

    observed, err := CalculateTotal(schematic)
    if err != nil {
        t.Fatalf("Error in function %v", err)
    }

    if expected != observed {
        t.Fatalf("Wanted: %v, Got: %v", expected, observed)
    }
}

func TestCases(t *testing.T) {
    cases := []test {
        {
            name:       "none",
            want:       0,
            schematic:  []string {
                ".....",
                ".123.",
                ".....",
            },
        },
        {
            name:       "after",
            want:       123,
            schematic:  []string {
                ".....",
                ".123*",
                ".....",
            },
        },
        {
            name:       "before",
            want:       123,
            schematic:  []string {
                ".....",
                "*123.",
                ".....",
            },
        },
        {
            name:       "top right",
            want:       123,
            schematic:  []string {
                "....*",
                ".123.",
                ".....",
            },
        },
        {
            name:       "top left",
            want:       123,
            schematic:  []string {
                "*....",
                ".123.",
                ".....",
            },
        },
        {
            name:       "bottom right",
            want:       123,
            schematic:  []string {
                ".....",
                ".123.",
                "....*",
            },
        },
        {
            name:       "bottom left",
            want:       123,
            schematic:  []string {
                ".....",
                ".123.",
                "*....",
            },
        },
        {
            name:       "top",
            want:       123,
            schematic:  []string {
                "..*..",
                ".123.",
                ".....",
            },
        },
        {
            name:       "bottom",
            want:       123,
            schematic:  []string {
                ".....",
                ".123.",
                "..*..",
            },
        },
    }
    
    for _, v := range cases {
        t.Run(v.name, func(t *testing.T) {
            result, err := CalculateTotal(v.schematic)
            if err != nil {
                t.Fatalf("Error: %v", err)
            }

            if result != v.want {
                t.Errorf("Want: %v Got: %v", v.want, result)
            }

        })
    }
}
