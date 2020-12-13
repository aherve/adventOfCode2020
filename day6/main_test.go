package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStr struct {
	Input    string
	Expected int
}

func TestAllScore(t *testing.T) {
	tests := []testStr{
		{
			Input:    "abc",
			Expected: 3,
		},
		{
			Input:    "a\nb\nc",
			Expected: 0,
		},
		{
			Input:    "ab\nac",
			Expected: 1,
		},
		{
			Input:    "a\na\na\na",
			Expected: 1,
		},
		{
			Input:    "b",
			Expected: 1,
		},
		{
			Input:    "a\naaa",
			Expected: 1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, allScore(test.Input))
	}
}
