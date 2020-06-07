package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldFindMaxSubarray(t *testing.T) {
	input := []float64{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	var actual, _ = maxSubarray(input)

	expected := []float64{18, 20, -7, 12}

	assert.Equal(t, expected, actual, "The two words should be the same.")
}

func TestShouldReturnInputWhenInputSizeIsOne(t *testing.T) {
	input := []float64{1}

	var actual, _ = maxSubarray(input)

	assert.Equal(t, input, actual)
}

func TestAllNegativeArray(t *testing.T) {
	input := []float64{-1, -2, -4. - 10}

	var actual, _ = maxSubarray(input)

	var expected []float64

	assert.Equal(t, expected, actual)
}
