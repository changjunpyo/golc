package problems

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ps := int(math.Sqrt(float64(i)))
		if ps*ps == i {
			dp[i] = 1
			continue
		}
		tmp := 12345678
		for j := ps; j >= 1; j-- {
			if tmp > dp[i-j*j] {
				tmp = dp[i-j*j]
			}
		}
		dp[i] = tmp + 1
	}
	return dp[n]
}
