package BadVersion

var Test []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func isBadVersion(n int) bool {
	return n > 3
}

func Optimal(a []int) int {
	l := 0
	r := len(a) - 1

	for l < r {
		mid := min((l + r) / 2)

		if isBadVersion(a[mid]) {
			if mid == 1 || mid-1 > 0 && !isBadVersion(mid) {
				return mid
			} else {
				r = mid
			}
		} else {
			l = mid + 1
		}

	}

	return l
}
