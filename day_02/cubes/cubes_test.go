package cubes

import "testing"

type test struct {
    input   string
    want    uint 
}

func TestMain(t *testing.T) {
    tests := []test {
        {
            input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
            want: 48,
        },
        {
            input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
            want: 12,
        },
        {
            input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
            want: 1560,
        },
        {
            input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
            want: 630,
        },
        {
            input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
            want: 36,
        },
    }

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got, err := Cube(tt.input); got != tt.want {
                if err != nil {
                    t.Error(err)
                }
				t.Errorf("Want(%d)  = %d; want %d", tt.want , got, tt.want)
			}
		})
	}
}
