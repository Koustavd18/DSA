package longSubStr

import "math"

var SubStr = "abcabcbd"

func BruteForce(s string) float64 {
	inputLen := len(s)
	if inputLen <= 1 {
		return float64(inputLen)
	}
	maxLength := 0.00
	for l := 0; l < inputLen; l++ {
		var seen = make(map[string]bool)
		currLength := 0.00
		for r := l; r < inputLen; r++ {
			curChar := s[r]
			if seen[string(curChar)] {
				break
			}
			seen[string(curChar)] = true
			currLength += 1
			maxLength = math.Max(currLength, maxLength)
		}

	}

	return maxLength
}

func Optimal(s string) int {
	inputLen := len(s)

	if inputLen <= 1 {
		return inputLen
	}

	l := 0
	r := 0
	ans := 0

	var seen = make(map[string]int)

	for r < inputLen {
		curChar := s[r]
		if i, ok := seen[string(curChar)]; ok {
			l = max(l, i+1)
		}
		seen[string(curChar)] = r
		ans = max(ans, r-l+1)
		r++
	}

	return ans
}
