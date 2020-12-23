package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestV2Mask(t *testing.T) {
	reader := newReaderV2()
	line := "mask = 000000000000000000000000000000X1001X"
	reader.updateMask(line)
	expected := []int64{26, 27, 58, 59}
	actual := reader.applyMask(42, true)
	for i := range actual {
		t.Logf("got %b, expected %b", actual[i], expected[i])
	}
	assert.Equal(t, expected, actual)

}
