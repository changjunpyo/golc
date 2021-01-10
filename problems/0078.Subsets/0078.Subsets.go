package problems

var ans [][]int

func backtracking(now, num []int, idx int) {
	tmp := make([]int, len(now))
	copy(tmp, now)
	ans = append(ans, tmp)
	if idx >= len(num) {
		return
	}
	for i := idx; i < len(num); i++ {
		now = append(now, num[i])
		backtracking(now, num, i+1)
		now = now[:len(now)-1]
	}
	return
}

func subsets(num []int) [][]int {
	ans = [][]int{}
	backtracking([]int{}, num, 0)
	return ans
}
