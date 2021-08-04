package problems

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for ; n > 1 ; n /= 3{
		if n % 3 != 0{
			return false
		}
	}
	return true
}
