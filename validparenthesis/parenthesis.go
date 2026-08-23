package validparenthesis

const Test = "{[()]}"

func IsValid(s string) bool {

	bracks := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	stack := []rune{}

	for _, parenthesis := range s {
		if _, ok := bracks[parenthesis]; ok {
			stack = append(stack, parenthesis)
		} else {
			if len(stack) == 0 {
				return false
			}

			lastOpen := stack[len(stack)-1]

			if bracks[lastOpen] != parenthesis {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
