package MissingNumber

import (
	"slices"
)

var Test []int = []int{4, 5, 2, 0, 1}

func BruteForce(a []int) int {

	sortedInput := slices.Sorted(slices.Values(a))

	for i, val := range sortedInput {
		expectedVal := i

		if expectedVal != val {
			return expectedVal
		}
	}

	return -1
}

func Gauss(a []int) int {

	n := len(a)

	expectedSum := n * ((n + 1) / 2)
	actualSum := 0

	for _, val := range a {
		actualSum += val
	}

	return expectedSum - actualSum
}
