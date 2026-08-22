package phonenum

var Test = "23"

func LetterCombination(digits string) []string {

	if len(digits) <= 0 {
		return []string{}
	}

	ans := []string{}
	digs2String := map[rune]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	Backtrack(digits, &ans, digs2String, "", 0)
	return ans
}

func Backtrack(digits string, ans *[]string, digs2String map[rune]string, cur string, idx int) {

	if len(digits) == len(cur) {
		combination := append([]string(nil), cur)
		*ans = append(*ans, combination...)
		return
	}

	curDig := digits[idx]
	curString := digs2String[rune(curDig)]

	for _, char := range curString {
		cur += string(char)

		Backtrack(digits, ans, digs2String, cur, idx+1)
		cur = cur[:len(cur)-1]
	}
}
