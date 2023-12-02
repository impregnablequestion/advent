package extract

import "testing"

func TestExtractBaseCase(t *testing.T) {
    expected := 67
    if observed := Extract("sixsrv2df4ur5seven"); observed != expected {
        t.Fatalf("Extract('sixsrv2df4ur5seven') = %v, want %v", observed, expected)
    }
}

func TestCasesFromPage(t *testing.T) {

	tests := []struct {
		name      string
		want      int
	}{
		{
            name:       "two1nine",
			want:       29,
		},
		{
			name:      "eightwothree",
			want:      83,
		},
		{
			name:      "abcone2threexyz",
			want:      13,
		},
		{
			name:      "xtwone3four",
			want:      24,
		},
		{
			name:      "4nineeightseven2",
			want:      42,
		},
		{
			name:      "zoneight234",
			want:      14,
		},
		{
			name:      "7pqrstsixteen",
			want:      76,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Extract(tt.name)
			if got != tt.want {
				t.Errorf(
					"Got %d, want %d",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestExtractOnlyOneDigit(t *testing.T) {
    expected := 77
    if observed := Extract("furseven"); observed != expected {
        t.Fatalf("Extract('furseven') = %v, want %v", observed, expected)
    }
}

func TestExtractEmptyString(t *testing.T) {
    expected := 0
    if observed := Extract(""); observed != expected {
        t.Fatalf("Extract('') = %v, want %v", observed, expected)
    }
}
