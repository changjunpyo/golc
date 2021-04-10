package problems

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	two_before, one_before, now := 1, 2, 3
	for i := 2; i < n; i++ {
		now = one_before + two_before
		one_before, two_before = now, one_before
	}
	return now
}
