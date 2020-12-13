package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStr struct {
	Input    string
	Expected int
}

func TestSeatID(t *testing.T) {
	tests := []testStr{
		{
			Input:    "FBFBBFFRLR",
			Expected: 357,
		},
		{
			Input:    "BFFFBBFRRR",
			Expected: 567,
		},
		{
			Input:    "FFFBBBFRRR",
			Expected: 119,
		},
		{
			Input:    "BBFFBBFRLL",
			Expected: 820,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, seatID(test.Input))
	}
}
