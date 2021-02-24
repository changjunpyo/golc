package problems

func digitSqureSum(n int) int {
	ans, digit := 0, 0
	for n > 0 {
		digit = n % 10
		ans += digit * digit
		n /= 10
	}
	return ans
}

// 1. using hash map time: O(log)
/*
func isHappy(n int) bool {
    number := make(map[int]bool)

    for _, exist := number[n]; !exist; _, exist = number[n]{
        number[n] = true
        n = digitSqureSum(n)
        if n == 1{
            return true
        }
    }
    return false
}
*/

// 2. find cycle by using floyd's cycle-finding algorithm
func isHappy(n int) bool {

	slow, fast := n, digitSqureSum(n)

	for fast != 1 && slow != fast {
		slow = digitSqureSum(slow)
		fast = digitSqureSum(digitSqureSum(fast))
	}

	return fast == 1
}
