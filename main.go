package main

import (
	"math"
)

func maxSubarray(input []float64) ([]float64, float64) {
	if len(input) == 1 {
		return input, input[0]
	} else {
		leftMaxSubarray, leftMax := maxSubarray(input[:len(input)/2])
		rightMaxSubarray, rightMax := maxSubarray(input[len(input)/2:])
		crossingMaxSubarray, crossingMax := findCrossingMaxSubarray(input)

		if leftMax > rightMax && leftMax > crossingMax {
			return leftMaxSubarray, leftMax
		} else if rightMax > leftMax && rightMax > crossingMax {
			return rightMaxSubarray, rightMax
		} else {
			return crossingMaxSubarray, crossingMax
		}
	}
}

func findCrossingMaxSubarray(input []float64) ([]float64, float64) {
	leftSum := math.Inf(-1)
	var sumAccumulator float64 = 0
	leftSumPosition := (len(input) / 2) - 1
	for i := (len(input) / 2) - 1; i >= 0; i-- {
		sumAccumulator += input[i]
		if sumAccumulator > leftSum {
			leftSum = sumAccumulator
			leftSumPosition = i
		}
	}

	rightSum := math.Inf(-1)
	sumAccumulator = 0
	rightSumPosition := len(input) / 2

	for i := len(input) / 2; i < len(input); i++ {
		sumAccumulator += input[i]
		if sumAccumulator > rightSum {
			rightSum = sumAccumulator
			rightSumPosition = i
		}
	}

	var nil []float64
	if leftSum+rightSum < 0 {
		return nil, 0
	}
	return input[leftSumPosition : rightSumPosition+1], leftSum + rightSum
}
