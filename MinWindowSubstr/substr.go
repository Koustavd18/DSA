package minwindowsubstr

var Test string = "sabyck"
var Target string = "abc"

// Takes 2 input first input is a string second input is target sub string
func SlidingWindow(s, t string) (res string) {

	distance := int(^uint(0) >> 1)

	i, j := 0, len(t)

	for i < j && j < len(s)+1 {

		if ContainsAllChars(s[i:j], t) {

			if j-i < distance {
				res = s[i:j] // j since s[i:j] s[j] is not a part of the sub str
				distance = j - i
			}
			i++
		} else {
			j++
		}
	}

	return
}

func ContainsAllChars(s, t string) bool {
	if len(s) < len(t) {
		return false
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := need[s[i]]; ok {
			have[s[i]]++
		}
	}

	for ch, count := range need {
		if have[ch] < count {
			return false
		}
	}

	return true
}

func Optimal(s, t string) (res string) {

	len1 := len(s)
	len2 := len(t)

	if len2 > len1 {
		return
	}

	hashPat := make(map[byte]int)
	hashStr := make(map[byte]int)

	for i := 0; i < len2; i++ {
		char := t[i]
		hashPat[char] += 1
	}

	count := 0
	left := 0
	startIdx := -1
	minLen := int(^uint(0) >> 1)

	for right := 0; right < len1; right++ {
		ch := s[right]
		hashStr[ch]++

		if ptCount, ok := hashPat[ch]; ok && hashStr[ch] <= ptCount {
			count += 1
		}

		if count == len2 {
			for hashPat[s[left]] == 0 || hashStr[s[left]] > hashPat[s[left]] {
				leftCh := s[left]
				if hashStr[leftCh] > 0 {
					hashStr[leftCh]--
				}
				left++
			}

			window := right - left + 1

			if minLen > window {
				startIdx = left
				minLen = window
			}
		}

	}
	if startIdx == -1 {
		return ""
	}

	return s[startIdx : startIdx+minLen]
}
