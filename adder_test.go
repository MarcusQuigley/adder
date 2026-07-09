package adder

import "testing"

var testCasesInt = []struct {
	description string
	a           int
	b           int
	expected    int
}{
	{
		"1 and 1",
		1,
		1,
		2,
	},
	{
		"17 and 0",
		17,
		0,
		17,
	},
	{
		"16 and 21",
		16,
		21,
		37,
	},
	{
		"0 and 0",
		0,
		0,
		0,
	},
}

var testCasesFloats = []struct {
	description string
	a           float64
	b           float64
	expected    float64
}{
	{
		"1.2 and 1",
		1.2,
		1,
		2.2,
	},
	{
		"0 and 0",
		0,
		0,
		0.0,
	},
}

func TestInts(t *testing.T) {
	for _, tc := range testCasesInt {
		t.Run(tc.description, func(t *testing.T) {
			if got := Add(tc.a, tc.b); got != tc.expected {
				t.Errorf("got %d. want %d", got, tc.expected)
			}

		})
	}
}

func TestFloatss(t *testing.T) {
	for _, tc := range testCasesFloats {
		t.Run(tc.description, func(t *testing.T) {
			if got := Add(tc.a, tc.b); got != tc.expected {
				t.Errorf("got %.2f want %.2f", got, tc.expected)
			}

		})
	}
}
