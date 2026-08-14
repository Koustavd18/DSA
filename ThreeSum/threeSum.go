package threesum

import "slices"

var Test = []int{-1, 0, 1, 2, -1, -4} //[[-1, -1, 2], [-1, 0, 1]]

func Optimal(nums []int) [][]int {

	numS := slices.Sorted(slices.Values(nums))

	var res [][]int

	for i := 0; i < len(numS)-2; i++ {

		if i > 0 && numS[i] == numS[i-1] {
			continue
		}

		left, right := i+1, len(numS)-1

		for left < right {
			sum := numS[i] + numS[left] + numS[right]

			if sum == 0 {
				res = append(res, []int{numS[i], numS[left], numS[right]})
				left++
				right--

				for left < right && numS[left] == numS[left-1] {
					left++
				}
				for left < right && numS[right] == numS[right+1] {
					right--
				}

			} else if sum < 0 {
				left++
			} else {
				right--
			}

		}
	}

	return res
}
