package CountPrimes

func BruteForce(n int) int {
	if n < 2 {
		return 0
	}
	count := 0

	for i := 2; i < n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break

			}
		}
		if isPrime {
			count++
		}
	}
	return count
}

// func Optimal( n int)int {
// 	if n
// }
