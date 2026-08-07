package Binary

import "strconv"

var Binary1 string = "1010"
var Binary2 string = "1011"

func Add(b1 string, b2 string) string {
	i := len(b1) - 1
	j := len(b2) - 1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 || carry == 1 {

		if i >= 0 {
			carry += int(b1[i]) - '0'
			i--
		}
		if j >= 0 {
			carry += int(b2[j]) - '0'
			j--
		}
		result = strconv.Itoa(carry%2) + result
		carry /= 2
	}

	return result
}
