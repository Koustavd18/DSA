package FirstAndLast

import "math"

var Test []int = []int{10, 11, 11, 11, 17, 18}

func BruteForce(a []int, t int) (int, int) {

	var firstSeen int = -1
	var lastSeen int = -1

	for i, val := range a {
		if val == t {
			firstSeen = i
			break
		}
	}

	lastIndex := len(a) - 1
	i := lastIndex

	for i >= 0 {
		if a[i] == t {
			lastSeen = i
			break
		}
		i--
	}

	return firstSeen, lastSeen
}

func Optimal(a []int, t int) (int, int) {

	firstSeen := findFirst(a, t)
	lastSeen := findLast(a, t)

	return firstSeen, lastSeen
}

func findFirst(a []int, t int) int {
	left, right := 0, len(a)-1

	for right >= left {
		mid := int(math.Floor((float64(right) + float64(left)) / 2))
		if a[mid] == t {
			if mid == 0 || a[mid-1] != t {
				return mid
			}
			right = mid - 1
		} else if a[mid] > t {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func findLast(a []int, t int) int {

	left, right := 0, len(a)-1

	for right >= left {
		mid := int(math.Floor((float64(right) + float64(left)) / 2))
		if a[mid] == t {
			if mid == len(a)-1 || a[mid+1] != t {
				return mid
			}
			left = mid + 1
		} else if a[mid] > t {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return -1
}
