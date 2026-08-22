package palindromepart

const Test = "aab"

func Backtrack(s string) [][]string {

	result := [][]string{}
	partitioner(s, &result, []string{})
	return result
}

func partitioner(s string, result *[][]string, cur []string) {

	if len(s) == 0 {
		temp := make([]string, len(cur))
		copy(temp, cur)
		*result = append(*result, temp)
		return
	}

	for i := 1; i <= len(s); i++ {
		curString := s[0:i]

		if isPalindrome(curString) {
			cur = append(cur, curString)
			partitioner(s[i:], result, cur)
			cur = cur[:len(cur)-1]
		}
	}

}

func isPalindrome(s string) bool {

	l := 0
	r := len(s) - 1

	for l < r {

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
